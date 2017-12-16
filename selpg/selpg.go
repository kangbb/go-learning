package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const MAX_INT = 1<<32 - 1

/*--------------stuct of the command options and parameters-----------*/
type selpg_args struct {
	start_page  int
	end_page    int
	page_len    int
	page_type   string
	print_dest  string
	in_filename string
}

/*----------------a gloabal var to get program name------------------*/
/*the program as the command name*/
var progname string

/* --------------------------the main function-----------------------*/
func main() {
	var sa selpg_args

	progname = os.Args[0]

	flag.IntVar(&sa.start_page, "s", -1, "-sstart_page")
	flag.IntVar(&sa.end_page, "e", -1, "-eend_page")
	flag.IntVar(&sa.page_len, "l", 72, "-llines_per_page")
	flag.StringVar(&sa.print_dest, "d", "", "-ddest")

	/*define -f and check it whether exist
	-f only useful for bool, default false*/
	exist_f := flag.Bool("f", false, "-f=true")

	/*analysis the command-line parameters and save in flag*/
	flag.Parse()

	/*if f = true, read file use '\f' delimited*/
	if *exist_f {
		sa.page_type = "f"
		sa.page_len = -1
	} else {
		sa.page_type = "l"
	}

	/*the filename is the only Not Flag parameters*/
	if flag.NArg() == 1 {
		sa.in_filename = flag.Arg(0)
	} else {
		sa.in_filename = ""
	}

	/*check the parameters if meet the requirement*/
	checkArgs(sa, flag.NArg())

	/*exec the command*/
	process_input(sa)
}

/*------------------------for usage function------------------------------*/
func usage() {
	fmt.Fprintf(os.Stderr,
		"USAGE: %s -s start_page -e end_page [ -f | -l lines_per_page ] [ -d dest ] [ in_filename ]\n",
		progname)
}

/*------------------------checkArgs()------------------------------*/
func checkArgs(sa selpg_args, NArg int) {
	if !(sa.start_page <= sa.end_page && sa.start_page >= 1) ||
		!(NArg == 1 || NArg == 0) ||
		(sa.page_type == "f" && sa.page_len != -1) ||
		(sa.page_type == "l" && sa.page_len <= 0) {
		usage()
		os.Exit(1)
	}
}

/*------------------------process_input()------------------------------*/
func process_input(sa selpg_args) {
	var fin *os.File
	var fout *os.File
	var line_ctr int
	var page_ctr int
	var err error

	/*set the input source*/
	if sa.in_filename == "" {
		fin = os.Stdin
	} else {
		fin, err = os.Open(sa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n",
				progname, sa.in_filename)
			fmt.Println(err)
			os.Exit(12)
		}
	}

	/*set the ouput destination*/
	if sa.print_dest == "" {
		fout = os.Stdout
	} else {
		s1 := "lp -d"
		s1 += sa.print_dest

		cmd := exec.Command("lp", "-d", sa.print_dest)
		stdout, err := cmd.StderrPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open pipe to \"%s\"\n",
				progname, s1)
			fmt.Println(err)
			os.Exit(13)
		}

		err = cmd.Start()
		if err != nil {
			fmt.Println(err)
		}

		content, err := ioutil.ReadAll(stdout)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(string(content))
	}

	/*begin one of two main loops based on page type*/
	if sa.page_type == "l" {
		line_ctr = 0
		page_ctr = 1

		rd := bufio.NewReader(fin)
		for {
			line, err := rd.ReadString('\n')
			if err != nil || err == io.EOF {
				break
			}

			line = strings.Replace(line, "\f", "", -1)
			line_ctr++
			if line_ctr > sa.page_len {
				page_ctr++
				line_ctr = 1
			}
			if page_ctr >= sa.start_page && page_ctr <= sa.end_page {
				fmt.Fprintf(fout, "%s", line)
			}
		}
	} else {
		page_ctr = 0
		rd := bufio.NewReader(fin)
		for {
			page, err := rd.ReadString('\f')
			if err != nil || err == io.EOF {
				/*if have no '\f'*/
				if err == io.EOF {
					//page += "\n"  //convenient to test
					if page_ctr >= sa.start_page && page_ctr <= sa.end_page {
						fmt.Fprintf(fout, "%s", page)
					}
				}
				break
			}
			//page += "\n"  //convenient to test
			page = strings.Replace(page, "\f", "", -1)
			page_ctr++
			if page_ctr >= sa.start_page && page_ctr <= sa.end_page {
				fmt.Fprintf(fout, "%s", page)
			}
		}
	}

	/*end main loop*/
	if page_ctr < sa.start_page {
		fmt.Fprintf(os.Stderr,
			"%s: start_page (%d) greater than total pages (%d), no output written\n",
			progname, sa.start_page, page_ctr)
	} else if page_ctr < sa.end_page {
		fmt.Fprintf(os.Stderr,
			"%s: end_page (%d) greater than total pages (%d), less output than expected\n",
			progname, sa.end_page, page_ctr)
	}
	fin.Close()
	fout.Close()
	fmt.Fprintf(os.Stderr, "%s: done\n", progname)
}

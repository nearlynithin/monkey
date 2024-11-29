package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/nearlynithin/monkey/lexer"
	"github.com/nearlynithin/monkey/parser"
)

const MONKEY_FACE = `
                                   '$$$$*$$*******************                               
                                '$$********************************'                          
                             *$**************************************$*                       
                          *$***********''''''''''*****'''''''''''''''**$$$                    
                       '$$***'''''***'''''''''''''''''''*$$####$*''''''*'**$'                 
                     '$****'''''''''''''''''''''''''''*$########$$*'  ''''''$                 
                     **''''''''''''''''''''''''''''''*$###########$$$*''''                    
                    '**'''''''''''''''''''''''''''''*$$#####$$$$$$$$$$$$''                    
                    '''''''''''''''''''''''''''''''*$$$$$$$$$$$$$$$$$$$$$**'                  
                   ''''''''''    '''''''''''''''''*$$$$$$$$$$$$$$$$$$$$$$$*'                  
                  '''''''''''''    ' '         '''*$$$$$$$$$$$$$$$$$$$$$$$**                  
                  ''''''''''                    ''$$$$$$$$$$$$$$$$$$$$$$$$$*'                 
                  '''''''''''''''               ''*$$$$$$$$$$$$$$$$$$$$$$$$$*'                
                 '''''''''''                    '*$$$$$$$$$$$$$$#$$$$$$$$$$#%%#               
                 '''''''''''                     '*$$$$$$$$$$$#@@@@@%$$$$$#%@%%$              
                ''''''''''                        '$$$$$$$$$$' $%@@@@%$$$#* '%%#              
                '''''''''                          ''$$$$$$$'  '%%%%%#$$$$* '##$              
                '''''''                              ''*$$$$*  *#%#$**$$$$$$$**'              
                '''''                                 ''*$$$*''''''''*$###$$''                
                '                                      ''*$$$*''''''*$$####$$'                
                '                                       ''*$$$$$$$$$$$$$####$$'               
                        '''                            '***$$$$$$$$$$$$$$$$$$$$*              
                        **$'                          '*****$$$$$$$$$$$$$$$$$$$$'             
                        '$#*                          '****$$$$$$$$$$$$$$$*''''               
                         '''                         '**$$$$$$$$$$$$$$$$$*'                   
                                                    ''********$$*$*$$$$$$*'                   
                                                     '''''*****''*********''  ''              
                                                     ''''    ''            ''''               
                                                     '''****************$$$*'''               
                                                     ''''''''''''''******''''*                
                          '                            ' ''         '''''''''                 
                          *'                                           ''''                   
                         *$#######$*                                     '                    
                         *$###########$##%%%##$*'                                             
                          $$$#$######################%%%##*'                                  
                          *$$''''*$$$$$$########$$$$$$$$$$$####%#$*'                          
                          *%%%%%%#$*'''  '''*$$$$$$$$$$$$$$$$$$$$$$$##$$**                    
                        $%%%%#$#$#%#%%%%%#$$'    '''''''*****************'                    
                      *%%##*'''$$$$$$$$##%%%%%%##$$*'''       '''''''''*'                     
                      #$'*$%%%%%%%%%%##$$#%##%%#########$$****'''''''                         
                     '*%#%%%%%%%%%%%%%%##$''*##%%#########$$$*****'''                         
             *%$*   *'#%#%%%%%%%%%%%#######* '$#%##$$$$$$$$#$$$***' '                         
             *%######%%%%%%%%%%%%%%########$*' '$##%#%%%#######$$*'***                        
             *%######%%%%%%%%%%%%%%######$'**$*'*$*'''*$%%#####$$***$*'                       
             *%######%%%%%%%%%%%%%##$$$*'''*$#$*' ''*'*$**'$###$'*****'                       
             *%####%%%%%%%%%%%%%##$$*'' '*$##*''**'   '$#$**''''******'                       
             *%##%%%%%%%%%%%%%%##**'  '*$#%#*'*$*'    '$####$''*$******                       
             *##%%%%%%%%%%%%%###$*''*$$###$**###$*   '*####$$**$$******'                      
             *#%%%%%%%%%%%%###$*' '$$$$$*''$$$$$#$''''*###$$***********'                      
`

const PROMPT=">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		
		if len(p.Errors()) != 0 {
			printParserErrors(out,p.Errors())
			continue
		}
	
		io.WriteString(out, program.String())
		io.WriteString(out,"\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out,MONKEY_FACE)
	io.WriteString(out,"\t\ti don't know what this means bruh\n")
	io.WriteString(out,"Parser errors :\n")
	for _,msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
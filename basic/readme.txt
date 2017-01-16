importing code as libraries
http://stackoverflow.com/questions/23695448/golang-run-all-go-files-within-current-directory-through-the-command-line-mul

Error in importing custom packages in Go Lang
http://stackoverflow.com/questions/25501875/error-in-importing-custom-packages-in-go-lang

You would need to make your function exportable with an uppercase for its name
func Fastget(...
Used as:
n:=libfastget.Fastget(url,4,filename)
The spec mentions: "Exported identifiers":
    An identifier may be exported to permit access to it from another package. An identifier is exported if both:
   		the first character of the identifier's name is a Unicode upper case letter (Unicode class "Lu"); and
   		the identifier is declared in the package block or it is a field name or method name.
    All other identifiers are not exported. 
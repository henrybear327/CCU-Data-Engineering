// regex_match example
#include <iostream>
#include <string>
#include <regex>
#include <string>

int main ()
{
  std::string str, regex;
  int repeat = 30;

  // for(int i = 0; i < repeat * 2; i++) // exact
  //   str += "a";

  for(int i = 0; i < repeat; i++) // partial
    str += "a";
  
  for(int i = 0; i < repeat; i++)
    regex += "a?";
  for(int i = 0; i < repeat; i++)
    regex += "a";
  
  std::cout << str << std::endl;
  std::cout << regex << std::endl;

  if (std::regex_match (str, std::regex(regex) ))
    std::cout << "string literal matched\n";

    /*
  const char cstr[] = "subject";
  std::string s ("subject");
  std::regex e ("(sub)(.*)");

  if (std::regex_match (s,e))
    std::cout << "string object matched\n";

  if ( std::regex_match ( s.begin(), s.end(), e ) )
    std::cout << "range matched\n";

  std::cmatch cm;    // same as std::match_results<const char*> cm;
  std::regex_match (cstr,cm,e);
  std::cout << "string literal with " << cm.size() << " matches\n";

  std::smatch sm;    // same as std::match_results<string::const_iterator> sm;
  std::regex_match (s,sm,e);
  std::cout << "string object with " << sm.size() << " matches\n";

  std::regex_match ( s.cbegin(), s.cend(), sm, e);
  std::cout << "range with " << sm.size() << " matches\n";

  // using explicit flags:
  std::regex_match ( cstr, cm, e, std::regex_constants::match_default );

  std::cout << "the matches were: ";
  for (unsigned i=0; i<cm.size(); ++i) {
    std::cout << "[" << cm[i] << "] ";
  }
    */
  std::cout << std::endl;

  return 0;
}
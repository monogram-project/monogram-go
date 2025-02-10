1 -> pop_file_versions;

extend_searchlist( 'auto', popautolist ) -> popautolist;
extend_searchlist( 'lib', popuseslist ) -> popuseslist;

load monogram.p

lvars f;
for f in ['01' '02' '03' '04' '05' '06'] do
    monogram(discin('../ex/' >< f >< '.mg')) ==>
endfor;

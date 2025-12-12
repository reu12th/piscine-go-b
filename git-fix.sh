#!/bin/bash

# Auto Commit Script for 01-edu Piscine
# It commits each file with its assigned message.

declare -A files

# ---------------------------
# QUEST 01: Shell & Git
# ---------------------------
files["my_file.sh"]="q1 t1"
files["createdir.sh"]="q1 t2"
files["printgroups.sh"]="q1 t3"
files["who-are-you.sh"]="q1 t4"
files["z1"]="q1 t5"
files["lookagain.sh"]="q1 t6"
files["mastertheoc.sh"]="q1 t7"
files["hello.sh"]="q1 t8"
files["organs.sh"]="q1 t9"
files["countfiles.sh"]="q1 t10"
files["git-fix.sh"]="q1 t11"

# ---------------------------
# QUEST 02: Basics
# ---------------------------
files["printalphabet/main.go"]="q2 t1"
files["printreversealphabet/main.go"]="q2 t2"
files["printdigits/main.go"]="q2 t3"
files["isnegative.go"]="q2 t4"
files["printcomb.go"]="q2 t5"
files["printcomb2.go"]="q2 t6"
files["printnbr.go"]="q2 t7"
files["printcombn.go"]="q2 t8"

# ---------------------------
# QUEST 03: Pointers & Arithmetic
# ---------------------------
files["pointone.go"]="q3 t1"
files["ultimatepointone.go"]="q3 t2"
files["divmod.go"]="q3 t3"
files["ultimatedivmod.go"]="q3 t4"
files["printstr.go"]="q3 t5"
files["strlen.go"]="q3 t6"
files["swap.go"]="q3 t7"
files["strrev.go"]="q3 t8"
files["basicatoi.go"]="q3 t9"
files["basicatoi2.go"]="q3 t10"
files["atoi.go"]="q3 t11"
files["sortintegertable.go"]="q3 t12"

# ---------------------------
# QUEST 04: Iteration & Recursion
# ---------------------------
files["iterativefactorial.go"]="q4 t1"
files["recursivefactorial.go"]="q4 t2"
files["iterativepower.go"]="q4 t3"
files["recursivepower.go"]="q4 t4"
files["fibonacci.go"]="q4 t5"
files["sqrt.go"]="q4 t6"
files["isprime.go"]="q4 t7"
files["findnextprime.go"]="q4 t8"
files["eightqueens.go"]="q4 t9"

# ---------------------------
# QUEST 05: Strings & Runes
# ---------------------------
files["firstrune.go"]="q5 t1"
files["nrune.go"]="q5 t2"
files["lastrune.go"]="q5 t3"
files["index.go"]="q5 t4"
files["compare.go"]="q5 t5"
files["toupper.go"]="q5 t6"
files["tolower.go"]="q5 t7"
files["capitalize.go"]="q5 t8"
files["trimatoi.go"]="q5 t9"
files["isalpha.go"]="q5 t10"
files["isnumeric.go"]="q5 t11"
files["islower.go"]="q5 t12"
files["isupper.go"]="q5 t13"
files["isprintable.go"]="q5 t14"
files["concat.go"]="q5 t15"
files["basicjoin.go"]="q5 t16"
files["join.go"]="q5 t17"
files["alphacount.go"]="q5 t18"
files["printnbrinorder.go"]="q5 t19"
files["printnbrbase.go"]="q5 t20"
files["atoibase.go"]="q5 t21"

# ---------------------------
# QUEST 06: Command Line Args
# ---------------------------
files["printprogramname/main.go"]="q6 t1"
files["printparams/main.go"]="q6 t2"
files["revparams/main.go"]="q6 t3"
files["sortparams/main.go"]="q6 t4"

# ---------------------------
# QUEST 07: Memory & Slices
# ---------------------------
files["appendrange.go"]="q7 t1"
files["makerange.go"]="q7 t2"
files["concatparams.go"]="q7 t3"
files["splitwhitespaces.go"]="q7 t4"
files["printwordstables.go"]="q7 t5"
files["convertbase.go"]="q7 t6"

# ---------------------------
# QUEST 08: Structs & File I/O
# ---------------------------
files["boolean/main.go"]="q8 t1"
files["displayfile/main.go"]="q8 t2"
files["cat/main.go"]="q8 t3"
files["zjmp/main.go"]="q8 t4"

# ---------------------------
# QUEST 09: Functional Programming
# ---------------------------
files["foreach.go"]="q9 t1"
files["map.go"]="q9 t2"
files["any.go"]="q9 t3"
files["countif.go"]="q9 t4"
files["issorted.go"]="q9 t5"
files["doop/main.go"]="q9 t6"


#######################################
# QUEST 11: Linked Lists (Updated)
#######################################
files["listpushback.go"]="q11 t1"
files["listpushfront.go"]="q11 t2"
files["listsize.go"]="q11 t3"
files["listlast.go"]="q11 t4"
files["listclear.go"]="q11 t5"
files["listat.go"]="q11 t6"
files["listreverse.go"]="q11 t7"
files["listforeach.go"]="q11 t8"
files["listforeachif.go"]="q11 t9"
files["listfind.go"]="q11 t10"
files["listremoveif.go"]="q11 t11"
files["listmerge.go"]="q11 t12"
files["listsort.go"]="q11 t13"
files["sortlistinsert.go"]="q11 t14"
files["sortedlistmerge.go"]="q11 t15"

#######################################
# QUEST 12: Binary Trees (Updated)
#######################################
files["btreeinsertdata.go"]="q12 t1"
files["btreeapplyinfix.go"]="q12 t2"
files["btreeapplyprefix.go"]="q12 t3"
files["btreeapplysuffix.go"]="q12 t4"
files["btreesearchitem.go"]="q12 t5"
files["btreelevelcount.go"]="q12 t6"
files["btreeisbinary.go"]="q12 t7"
files["btreeapplybylevel.go"]="q12 t8"
files["btreemax.go"]="q12 t9"
files["btreemin.go"]="q12 t10"
files["btreetransplant.go"]="q12 t11"
files["btreedeletenode.go"]="q12 t12"

echo "---- Starting 01-edu Auto Commit + Push ----"

for file in "${!files[@]}"; do
    if [[ -f "$file" ]]; then
        echo "Processing: $file"

        git add "$file"
        git commit -m "${files[$file]}"
        git push

        echo "✔ Done: $file => ${files[$file]}"
        echo "-------------------------------------------"
    else
        echo "⚠ Skipped (not found): $file"
    fi
done

echo "---- All files processed ----"

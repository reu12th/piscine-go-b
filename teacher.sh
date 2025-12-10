interviewn=$(grep -H "license" interviews/* | grep "\" "| cut -f1 -d ":"| rev | cut -f1 -d "-" | rev)

export interviewnum="$interviewn"

echo "$interviewnum"

cat "interviews/interview-$interviewn"

echo $MAIN_SUSPECT
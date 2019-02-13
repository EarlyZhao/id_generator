#!/bin/bash
lua_script_file='/tmp/wrk_id_generator.lua'
wrk_id_generator_log_file='/tmp/wrk_id_generator_result_log.txt'

lua_txt="fileName = '$wrk_id_generator_log_file'
f = io.open(fileName, 'a')
response = function(status, headers, body)
  if status == 200 then
    f:write(body..'\n')
  end
end"

echo $lua_txt > $lua_script_file

wrk_http_url='http://127.0.0.1:1314/unique_ids/test'
wrk_connections=50
wrk_duration=10

wrk -c $wrk_connections -t $wrk_connections -d $wrk_duration  $wrk_http_url -s $lua_script_file

lines_count=`cat $wrk_id_generator_log_file | wc -l`
uniq_count=`cat $wrk_id_generator_log_file | uniq | wc -l`

echo ""
echo "check ID unique:"
echo "total lines: $lines_count"
echo "unique lines: $uniq_count"

same=$(($lines_count-$uniq_count))
if [ "$same" == "0" ]; then
  echo -e "\033[32m No repeated ID \033[0m"
  rm $lua_script_file
  rm $wrk_id_generator_log_file
else
  echo -e "\033[31m ID not unique, check $wrk_id_generator_log_file \033[0m"
fi

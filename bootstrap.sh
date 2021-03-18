# Author: d1y<chenhonzhou@gmail.com>
# create date 2021/03/09

mvn clean package

# TODO
docker rm -f clean-gnote
docker build -t clean-gnote .
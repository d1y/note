# Author: d1y<chenhonzhou@gmail.com>
# create date 2021/03/09

FROM mayan31370/openjdk-alpine-with-chinese-timezone

COPY target/note.app.jar .

CMD ["java", "-jar", "note.app.jar"]

EXPOSE 8080
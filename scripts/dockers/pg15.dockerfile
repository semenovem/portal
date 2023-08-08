FROM postgres:15.3
RUN apt-get clean && apt-get update && apt-get install -y locales
RUN localedef -i ru_RU -c -f UTF-8 -A /usr/share/locale/locale.alias ru_RU.UTF-8
ENV LANG ru_RU.utf8


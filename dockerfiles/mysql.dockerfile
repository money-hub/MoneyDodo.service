FROM mysql:5.7
# no password
ENV MYSQL_ALLOW_EMPTY_PASSWORD yes

# put the file to container
COPY dockerfiles/mysql_setup.sh /mysql/mysql_setup.sh
COPY db/priviledges.sql /mysql/priviledges.sql
COPY db/data.sql /mysql/data.sql
# command
CMD ["sh", "/mysql/mysql_setup.sh"]
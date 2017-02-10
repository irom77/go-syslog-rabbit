FROM logstash:5

# Add your logstash plugins setup here
# RUN logstash-plugin install logstash-filter-de_dot
# Example: RUN logstash-plugin install logstash-filter-json
RUN logstash-plugin install logstash-output-rabbitmq logstash-output-exec logstash-filter-json_encode

COPY ./config/* /etc/logstash/conf.d/
#COPY ./opt/* /opt/
#COPY ./exec/* /usr/local/bin/

CMD ["logstash", "-f", "/etc/logstash/conf.d/"]
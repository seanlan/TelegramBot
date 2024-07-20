CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go && \
rsync --progress main tx-tele-bot:/data/telegram-bot/main.new && \
ssh tx-tele-bot "supervisorctl stop telegram-bot && \
                \cp /data/telegram-bot/main /data/telegram-bot/main.last && \
                \cp /data/telegram-bot/main.new /data/telegram-bot/main && \
                supervisorctl start telegram-bot" && \
rsync --progress main tx-tele-bot-2:/data/telegram-bot/main.new && \
ssh tx-tele-bot-2 "supervisorctl stop telegram-bot && \
                \cp /data/telegram-bot/main /data/telegram-bot/main.last && \
                \cp /data/telegram-bot/main.new /data/telegram-bot/main && \
                supervisorctl start telegram-bot" && \
rm -rf main
npm run build:prod && \
scp -r dist/* tx-tele-bot:/data/telegram-bot/www && \
scp -r dist/* tx-tele-bot-2:/data/telegram-bot/www

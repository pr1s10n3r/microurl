# Micro URL

Small URL shortener project. I did it because I was boring so do not expect any production-ready code.

## Database Migration

You will need [flyway](https://flywaydb.org/). Once installed, go to **project directory** and execute this command:

```sh
flyway -configFiles=database/flyway.conf migrate
```

This command will run every SQL migration into `database/migrations` directory.

## Author

Hi! My name is Alvaro and I write code. You can check my [website](https://astagg.me); I write about programming there.
See you around!
FROM php:8.3-fpm

RUN apt-get update && apt-get install -y zlib1g-dev g++ git libicu-dev zip libzip-dev zip libssl-dev librabbitmq-dev gcc make autoconf libc-dev pkg-config \
    && docker-php-ext-install intl opcache pdo pdo_mysql \
    && pecl install apcu \
    && docker-php-ext-enable apcu \
    && pecl install amqp \
    && docker-php-ext-enable amqp \
    && docker-php-ext-configure zip \
    && docker-php-ext-install zip

RUN apt-get install -y supervisor
COPY email-service/supervisor/ /etc/supervisor/conf.d/
CMD ["supervisord", "-n",  "-c" , "/etc/supervisor/supervisord.conf"]
RUN echo 'cao'
WORKDIR /var/www/project

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer


RUN curl -sS https://get.symfony.com/cli/installer | bash


<?php

namespace App\Service\ResponseSagaItemResolver;

use App\Message\ResponseSagaItemMessage;

interface ResponseSagaItemResolverInterface
{
    public function resolve(ResponseSagaItemMessage $responseSagaItemMessage): void;

    public static function getResolverType(): string;
}
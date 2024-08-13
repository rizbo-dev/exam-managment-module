<?php

namespace App\Service\SagaItemDispatcher;

use App\Entity\SagaItem;

interface SagaItemDispatcherInterface
{
    public function dispatch(SagaItem $sagaItem): void;

    public static function getDispatcherType(): string;
}
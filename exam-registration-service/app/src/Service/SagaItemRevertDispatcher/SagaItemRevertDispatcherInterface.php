<?php

namespace App\Service\SagaItemRevertDispatcher;

use App\Entity\SagaItem;

interface SagaItemRevertDispatcherInterface
{
    public function revert(SagaItem $sagaItem): void;
    public static function getDispatcherType(): string;

}
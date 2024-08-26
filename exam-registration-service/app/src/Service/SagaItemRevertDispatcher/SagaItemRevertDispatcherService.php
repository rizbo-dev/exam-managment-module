<?php

namespace App\Service\SagaItemRevertDispatcher;

use App\Entity\SagaItem;

class SagaItemRevertDispatcherService
{
    /**
     * @var array<string, SagaItemRevertDispatcherInterface>
     */
    private array $dispatchers;

    public function __construct(
        iterable $dispatchers,

    )
    {
        $this->dispatchers = $dispatchers instanceof \Traversable ? iterator_to_array($dispatchers) : $dispatchers;
    }

    public function revert(SagaItem $sagaItem): void
    {
        if (!isset($this->dispatchers[$sagaItem->getSagaType()])) {
            return;
        }

        $this->dispatchers[$sagaItem->getSagaType()]->revert($sagaItem);
    }
}
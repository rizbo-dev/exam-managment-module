<?php

namespace App\Service\SagaItemDispatcher;

use App\Entity\SagaItem;
use Doctrine\ORM\EntityManagerInterface;

class SagaItemDispatcherService
{
    /**
     * @var array<string, SagaItemDispatcherInterface>
     */
    private array $dispatchers;

    public function __construct(
        iterable $dispatchers,
        private readonly EntityManagerInterface $entityManager
    )
    {
        $this->dispatchers = $dispatchers instanceof \Traversable ? iterator_to_array($dispatchers) : $dispatchers;
    }

    public function dispatch(SagaItem $sagaItem): void
    {
        if (!isset($this->dispatchers[$sagaItem->getSagaType()])) {
            throw new \RuntimeException('Unhandled type');
        }

        $sagaItem
            ->setStartedAt(new \DateTimeImmutable())
            ->setStatus(SagaItem::IN_PROGRESS);

        $this->entityManager->flush();

        $this->dispatchers[$sagaItem->getSagaType()]->dispatch($sagaItem);
    }
}
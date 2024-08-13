<?php

namespace App\MessageHandler;

use App\Message\ResponseSagaItemMessage;
use App\Service\ResponseSagaItemResolver\ResponseSagaItemResolverInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
readonly class ResponseSagaItemMessageHandler
{
    /**
     * @var array<string, ResponseSagaItemResolverInterface>
     */
    private array $resolvers;

    public function __construct(
        iterable $resolvers
    )
    {
        $this->resolvers = $resolvers instanceof \Traversable ? iterator_to_array($resolvers) : $resolvers;
    }

    public function __invoke(ResponseSagaItemMessage $responseSagaItemMessage): void
    {
        if (!isset($this->resolvers[$responseSagaItemMessage->getSagaType()])) {
            throw new \RuntimeException('Unhandled type');
        }

        $this->resolvers[$responseSagaItemMessage->getSagaType()]->resolve($responseSagaItemMessage);
    }
}
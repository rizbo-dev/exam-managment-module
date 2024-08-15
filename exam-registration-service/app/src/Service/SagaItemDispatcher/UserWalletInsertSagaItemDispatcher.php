<?php

namespace App\Service\SagaItemDispatcher;

use App\Entity\SagaItem;
use App\Message\UserWalletInsertMessage;
use App\Service\SagaItemService;
use Symfony\Component\Messenger\MessageBusInterface;

readonly class UserWalletInsertSagaItemDispatcher implements SagaItemDispatcherInterface
{
    public function __construct(
        private MessageBusInterface $messageBus
    )
    {
    }

    public function dispatch(SagaItem $sagaItem): void
    {
        $message = new UserWalletInsertMessage();

        $message
            ->setStudentId($sagaItem->getExamRegistration()->getStudentId())
            ->setExamRegistrationId($sagaItem->getExamRegistration()->getId())
            ->setAmount($this->getAmountForInsert($sagaItem));


        $this->messageBus->dispatch($message);
    }

    public static function getDispatcherType(): string
    {
        return SagaItem::USER_WALLET_INSERT_SAGA_ITEM_TYPE;
    }

    private function getAmountForInsert(SagaItem $sagaItem): int
    {
        $validationSaga = SagaItemService::getUserWalletValidationSagaItem($sagaItem->getExamRegistration());

        return $validationSaga->getReturnedPayload()['examCreditValue'];
    }
}
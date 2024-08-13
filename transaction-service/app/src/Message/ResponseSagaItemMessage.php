<?php

namespace App\Message;

class ResponseSagaItemMessage
{
    public const USER_WALLET_VALIDATION_SAGA_ITEM_TYPE = 'userWalletValidationSagaItem';

    private string $status;
    private string $sagaType;
    private int $examRegistrationId;

    private array $payload;

    public function __construct(string $status, string $sagaType, int $examRegistrationId, array $payload)
    {
        $this->status = $status;
        $this->sagaType = $sagaType;
        $this->examRegistrationId = $examRegistrationId;
        $this->payload = $payload;
    }

    public function getStatus(): string
    {
        return $this->status;
    }

    public function setStatus(string $status): void
    {
        $this->status = $status;
    }

    public function getSagaType(): string
    {
        return $this->sagaType;
    }

    public function setSagaType(string $sagaType): void
    {
        $this->sagaType = $sagaType;
    }

    public function getExamRegistrationId(): int
    {
        return $this->examRegistrationId;
    }

    public function setExamRegistrationId(int $examRegistrationId): void
    {
        $this->examRegistrationId = $examRegistrationId;
    }

    public function getPayload(): array
    {
        return $this->payload;
    }

    public function setPayload(array $payload): ResponseSagaItemMessage
    {
        $this->payload = $payload;
        return $this;
    }
}
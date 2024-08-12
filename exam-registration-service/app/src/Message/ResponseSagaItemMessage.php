<?php

namespace App\Message;

class ResponseSagaItemMessage
{
    public const USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE = 'userClassVerificationSagaItem';

    private string $status;
    private string $sagaType;
    private int $examRegistrationId;

    public function __construct(string $status, string $sagaType, int $examRegistrationId)
    {
        $this->status = $status;
        $this->sagaType = $sagaType;
        $this->examRegistrationId = $examRegistrationId;
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
}
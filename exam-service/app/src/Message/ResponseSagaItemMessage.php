<?php

namespace App\Message;

class ResponseSagaItemMessage
{
    private string $status;
    private string $sagaType;
    private int $examRegistrationId;
    private array $payload;

    public function getStatus(): string
    {
        return $this->status;
    }

    public function setStatus(string $status): ResponseSagaItemMessage
    {
        $this->status = $status;
        return $this;
    }

    public function getSagaType(): string
    {
        return $this->sagaType;
    }

    public function setSagaType(string $sagaType): ResponseSagaItemMessage
    {
        $this->sagaType = $sagaType;
        return $this;
    }

    public function getExamRegistrationId(): int
    {
        return $this->examRegistrationId;
    }

    public function setExamRegistrationId(int $examRegistrationId): ResponseSagaItemMessage
    {
        $this->examRegistrationId = $examRegistrationId;
        return $this;
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
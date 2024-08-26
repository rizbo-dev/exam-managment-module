<?php

namespace App\Message;

class UserUpdatedMessage
{
    private int $userId;

    private string $actionType;

    public function getUserId(): int
    {
        return $this->userId;
    }

    public function setUserId(int $userId): UserUpdatedMessage
    {
        $this->userId = $userId;
        return $this;
    }

    public function getActionType(): string
    {
        return $this->actionType;
    }

    public function setActionType(string $actionType): UserUpdatedMessage
    {
        $this->actionType = $actionType;
        return $this;
    }
}
<?php

namespace App\Message;

class UserWalletRevertMessage
{
    private int $transactionId;

    public function getTransactionId(): int
    {
        return $this->transactionId;
    }

    public function setTransactionId(int $transactionId): UserWalletRevertMessage
    {
        $this->transactionId = $transactionId;
        return $this;
    }
}
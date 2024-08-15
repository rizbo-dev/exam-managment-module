<?php

namespace App\Message;

class UserWalletInsertMessage
{
    private int $studentId;
    private int $amount;
    private int $examRegistrationId;

    public function getStudentId(): int
    {
        return $this->studentId;
    }

    public function setStudentId(int $studentId): UserWalletInsertMessage
    {
        $this->studentId = $studentId;
        return $this;
    }

    public function getAmount(): int
    {
        return $this->amount;
    }

    public function setAmount(int $amount): UserWalletInsertMessage
    {
        $this->amount = $amount;
        return $this;
    }

    public function getExamRegistrationId(): int
    {
        return $this->examRegistrationId;
    }

    public function setExamRegistrationId(int $examRegistrationId): UserWalletInsertMessage
    {
        $this->examRegistrationId = $examRegistrationId;
        return $this;
    }
}
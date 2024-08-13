<?php

namespace App\Message;

class UserWalletValidationMessage
{
    private int $studentId;
    private int $examId;
    private int $examRegistrationId;

    public function __construct(int $studentId, int $examId, int $examRegistrationId)
    {
        $this->studentId = $studentId;
        $this->examId = $examId;
        $this->examRegistrationId = $examRegistrationId;
    }

    public function getStudentId(): int
    {
        return $this->studentId;
    }

    public function setStudentId(int $studentId): UserWalletValidationMessage
    {
        $this->studentId = $studentId;
        return $this;
    }

    public function getExamId(): int
    {
        return $this->examId;
    }

    public function setExamId(int $examId): UserWalletValidationMessage
    {
        $this->examId = $examId;
        return $this;
    }

    public function getExamRegistrationId(): int
    {
        return $this->examRegistrationId;
    }

    public function setExamRegistrationId(int $examRegistrationId): UserWalletValidationMessage
    {
        $this->examRegistrationId = $examRegistrationId;
        return $this;
    }
}
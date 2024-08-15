<?php

namespace App\Message;

class ExamRegistrationMessage
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

    public function setStudentId(int $studentId): ExamRegistrationMessage
    {
        $this->studentId = $studentId;
        return $this;
    }

    public function getExamId(): int
    {
        return $this->examId;
    }

    public function setExamId(int $examId): ExamRegistrationMessage
    {
        $this->examId = $examId;
        return $this;
    }

    public function getExamRegistrationId(): int
    {
        return $this->examRegistrationId;
    }

    public function setExamRegistrationId(int $examRegistrationId): ExamRegistrationMessage
    {
        $this->examRegistrationId = $examRegistrationId;
        return $this;
    }
}
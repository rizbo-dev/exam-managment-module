<?php

namespace App\Message;

class UserClassVerificationMessage
{
    private int $studentId;
    private int $examId;
    private int $examRegistrationId;
    private int $courseId;

    public function __construct(int $studentId, int $examId, int $examRegistrationId, int $courseId)
    {
        $this->studentId = $studentId;
        $this->examId = $examId;
        $this->examRegistrationId = $examRegistrationId;
        $this->courseId = $courseId;
    }

    public function getStudentId(): int
    {
        return $this->studentId;
    }

    public function setStudentId(int $studentId): void
    {
        $this->studentId = $studentId;
    }

    public function getExamId(): int
    {
        return $this->examId;
    }

    public function setExamId(int $examId): void
    {
        $this->examId = $examId;
    }

    public function getExamRegistrationId(): int
    {
        return $this->examRegistrationId;
    }

    public function setExamRegistrationId(int $examRegistrationId): void
    {
        $this->examRegistrationId = $examRegistrationId;
    }

    public function getCourseId(): int
    {
        return $this->courseId;
    }

    public function setCourseId(int $courseId): UserClassVerificationMessage
    {
        $this->courseId = $courseId;
        return $this;
    }
}
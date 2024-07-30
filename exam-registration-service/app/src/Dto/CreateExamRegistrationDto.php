<?php

namespace App\Dto;

class CreateExamRegistrationDto
{
    private int $studentId;

    private int $examId;

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
}
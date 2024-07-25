<?php

namespace App\Entity;

use App\Repository\ExamStudentRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ExamStudentRepository::class)]
class ExamStudent
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column]
    private ?int $studentId = null;

    #[ORM\ManyToOne(inversedBy: 'examStudents')]
    private ?Exam $exam = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getStudentId(): ?int
    {
        return $this->studentId;
    }

    public function setStudentId(int $studentId): static
    {
        $this->studentId = $studentId;

        return $this;
    }

    public function getExam(): ?Exam
    {
        return $this->exam;
    }

    public function setExam(?Exam $exam): static
    {
        $this->exam = $exam;

        return $this;
    }
}

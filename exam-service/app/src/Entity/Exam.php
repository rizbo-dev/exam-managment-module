<?php

namespace App\Entity;

use ApiPlatform\Metadata\ApiResource;
use App\Repository\ExamRepository;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ExamRepository::class)]
#[ApiResource]
class Exam
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column]
    private ?int $courseId = null;

    #[ORM\ManyToOne(inversedBy: 'exams')]
    private ?ExaminationPeriod $examinationPeriod = null;

    #[ORM\Column]
    private ?int $maxStudentEntry = null;

    #[ORM\Column]
    private ?float $cost = null;

    /**
     * @var Collection<int, ExamStudent>
     */
    #[ORM\OneToMany(targetEntity: ExamStudent::class, mappedBy: 'exam')]
    private Collection $examStudents;

    public function __construct()
    {
        $this->examStudents = new ArrayCollection();
    }

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getCourseId(): ?int
    {
        return $this->courseId;
    }

    public function setCourseId(int $courseId): static
    {
        $this->courseId = $courseId;

        return $this;
    }

    public function getExaminationPeriod(): ?ExaminationPeriod
    {
        return $this->examinationPeriod;
    }

    public function setExaminationPeriod(?ExaminationPeriod $examinationPeriod): static
    {
        $this->examinationPeriod = $examinationPeriod;

        return $this;
    }

    public function getMaxStudentEntry(): ?int
    {
        return $this->maxStudentEntry;
    }

    public function setMaxStudentEntry(int $maxStudentEntry): static
    {
        $this->maxStudentEntry = $maxStudentEntry;

        return $this;
    }

    public function getCost(): ?float
    {
        return $this->cost;
    }

    public function setCost(float $cost): static
    {
        $this->cost = $cost;

        return $this;
    }

    /**
     * @return Collection<int, ExamStudent>
     */
    public function getExamStudents(): Collection
    {
        return $this->examStudents;
    }

    public function addExamStudent(ExamStudent $examStudent): static
    {
        if (!$this->examStudents->contains($examStudent)) {
            $this->examStudents->add($examStudent);
            $examStudent->setExam($this);
        }

        return $this;
    }

    public function removeExamStudent(ExamStudent $examStudent): static
    {
        if ($this->examStudents->removeElement($examStudent)) {
            // set the owning side to null (unless already changed)
            if ($examStudent->getExam() === $this) {
                $examStudent->setExam(null);
            }
        }

        return $this;
    }
}

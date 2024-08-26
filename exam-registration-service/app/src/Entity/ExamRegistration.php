<?php

namespace App\Entity;

use App\Repository\ExamRegistrationRepository;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ExamRegistrationRepository::class)]
class ExamRegistration
{
    public const FINISHED_STATUS = 'finished';
    public const IN_PROGRESS_STATUS = 'in_progress';
    public const INITIALIZED_STATUS = 'initialized';

    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column]
    private ?int $studentId = null;

    #[ORM\Column]
    private ?int $examId = null;

    #[ORM\Column]
    private ?int $courseId = null;
    #[ORM\Column(length: 255)]
    private ?string $status = self::INITIALIZED_STATUS;

    /**
     * @var Collection<int, SagaItem>
     */
    #[ORM\OneToMany(targetEntity: SagaItem::class, mappedBy: 'examRegistration')]
    private Collection $sagaItems;

    public function __construct()
    {
        $this->sagaItems = new ArrayCollection();
    }

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

    public function getExamId(): ?int
    {
        return $this->examId;
    }

    public function setExamId(int $examId): static
    {
        $this->examId = $examId;

        return $this;
    }

    public function getStatus(): ?string
    {
        return $this->status;
    }

    public function setStatus(string $status): static
    {
        $this->status = $status;

        return $this;
    }

    /**
     * @return Collection<int, SagaItem>
     */
    public function getSagaItems(): Collection
    {
        return $this->sagaItems;
    }

    public function addSagaItem(SagaItem $sagaItem): static
    {
        if (!$this->sagaItems->contains($sagaItem)) {
            $this->sagaItems->add($sagaItem);
            $sagaItem->setExamRegistration($this);
        }

        return $this;
    }

    public function removeSagaItem(SagaItem $sagaItem): static
    {
        if ($this->sagaItems->removeElement($sagaItem)) {
            // set the owning side to null (unless already changed)
            if ($sagaItem->getExamRegistration() === $this) {
                $sagaItem->setExamRegistration(null);
            }
        }

        return $this;
    }

    public function getCourseId(): ?int
    {
        return $this->courseId;
    }

    public function setCourseId(?int $courseId): ExamRegistration
    {
        $this->courseId = $courseId;
        return $this;
    }
}

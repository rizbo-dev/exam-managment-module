<?php

namespace App\Entity;

use App\Repository\SagaItemRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: SagaItemRepository::class)]
class SagaItem
{
    public const INITIALIZED_STATUS = 'initialized';
    public const IN_PROGRESS = 'in_progress';
    public const FINISHED = 'finished';

    public const USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE = 'userClassVerificationSagaItem';
    public const USER_WALLET_VALIDATION_SAGA_ITEM_TYPE = 'userWalletValidationSagaItem';

    public const ITEMS = [
        [
            'sagaType' => self::USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE,
            'type' => 'validation',
            'executionOrder' => 1
        ],
        [
            'sagaType' => self::USER_WALLET_VALIDATION_SAGA_ITEM_TYPE,
            'type' => 'validation',
            'executionOrder' => 2
        ],
        [
            'sagaType' => 'userWalletInsertSagaItem',
            'type' => 'modification',
            'executionOrder' => 3
        ],
        [
            'sagaType' => 'examValidationSagaItem',
            'type' => 'validation',
            'executionOrder' => 4
        ],
        [
            'sagaType' => 'examRegistrationSagaItem',
            'type' => 'validation',
            'executionOrder' => 5
        ],
    ];

    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column(length: 255)]
    private ?string $type = null;

    #[ORM\Column(nullable: true)]
    private ?\DateTimeImmutable $startedAt = null;

    #[ORM\Column(nullable: true)]
    private ?\DateTimeImmutable $finishedAt = null;

    #[ORM\Column(length: 255)]
    private ?string $status = self::INITIALIZED_STATUS;

    #[ORM\Column]
    private ?int $executionOrder = null;

    #[ORM\Column(length: 255)]
    private ?string $sagaType = null;

    #[ORM\ManyToOne(inversedBy: 'sagaItems')]
    #[ORM\JoinColumn(nullable: false)]
    private ?ExamRegistration $examRegistration = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getType(): ?string
    {
        return $this->type;
    }

    public function setType(string $type): static
    {
        $this->type = $type;

        return $this;
    }

    public function getStartedAt(): ?\DateTimeImmutable
    {
        return $this->startedAt;
    }

    public function setStartedAt(\DateTimeImmutable $startedAt): static
    {
        $this->startedAt = $startedAt;

        return $this;
    }

    public function getFinishedAt(): ?\DateTimeImmutable
    {
        return $this->finishedAt;
    }

    public function setFinishedAt(?\DateTimeImmutable $finishedAt): static
    {
        $this->finishedAt = $finishedAt;

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

    public function getExecutionOrder(): ?int
    {
        return $this->executionOrder;
    }

    public function setExecutionOrder(int $executionOrder): static
    {
        $this->executionOrder = $executionOrder;

        return $this;
    }

    public function getSagaType(): ?string
    {
        return $this->sagaType;
    }

    public function setSagaType(string $sagaType): static
    {
        $this->sagaType = $sagaType;

        return $this;
    }

    public function getExamRegistration(): ?ExamRegistration
    {
        return $this->examRegistration;
    }

    public function setExamRegistration(?ExamRegistration $examRegistration): static
    {
        $this->examRegistration = $examRegistration;

        return $this;
    }
}

<?php

namespace App\EntityListener;

use App\Entity\Student;
use App\Message\StudentUpdateMessage;
use Doctrine\Bundle\DoctrineBundle\Attribute\AsEntityListener;
use Doctrine\ORM\Event\PostPersistEventArgs;
use Doctrine\ORM\Event\PostRemoveEventArgs;
use Doctrine\ORM\Event\PostUpdateEventArgs;
use Symfony\Component\Messenger\MessageBusInterface;

#[AsEntityListener]
readonly class StudentEntityListener
{
    public function __construct(
        private MessageBusInterface $messageBus
    )
    {
    }

    public function postPersist(Student $student, PostPersistEventArgs $args): void
    {
        $this->messageBus->dispatch(new StudentUpdateMessage($student->getId()));
    }

    public function postUpdate(Student $student, PostUpdateEventArgs $args): void
    {
        $this->messageBus->dispatch(new StudentUpdateMessage($student->getId()));
    }
}
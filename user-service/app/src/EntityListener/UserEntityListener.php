<?php

namespace App\EntityListener;

use App\Entity\User;
use App\Message\UserUpdatedMessage;
use Doctrine\Bundle\DoctrineBundle\Attribute\AsEntityListener;
use Doctrine\ORM\Event\PostPersistEventArgs;
use Doctrine\ORM\Event\PostUpdateEventArgs;
use Symfony\Component\Messenger\MessageBusInterface;

#[AsEntityListener]
class UserEntityListener
{
    public function __construct(
        private MessageBusInterface $messageBus
    )
    {
    }

    public function postPersist(User $user, PostPersistEventArgs $args)
    {
        $message = new UserUpdatedMessage();
        $message
            ->setUserId($user->getId())
            ->setActionType('create');

        $this->messageBus->dispatch($message);
    }

    public function postUpdate(User $user, PostUpdateEventArgs $args)
    {
        $message = new UserUpdatedMessage();
        $message
            ->setUserId($user->getId())
            ->setActionType('update');

        $this->messageBus->dispatch($message);
    }
}
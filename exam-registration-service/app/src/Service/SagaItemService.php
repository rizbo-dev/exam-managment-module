<?php

namespace App\Service;

use App\Entity\ExamRegistration;
use App\Entity\SagaItem;
use Doctrine\Common\Collections\Collection;

class SagaItemService
{
    /**
     * @param Collection<int, SagaItem> $sagaItems
     * @return SagaItem|null
     */
    public static function getNextItemForExecution(Collection $sagaItems): ?SagaItem
    {
        foreach ($sagaItems as $sagaItem) {
            if ($sagaItem->getStatus() === SagaItem::INITIALIZED_STATUS) {
                return $sagaItem;
            }
        }

        return null;
    }

    public static function getUserClassVerificationSagaItem(ExamRegistration $examRegistration): ?SagaItem
    {
        foreach ($examRegistration->getSagaItems() as $sagaItem) {
            if ($sagaItem->getSagaType() === SagaItem::USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE) {
                return $sagaItem;
            }
        }

        return null;
    }
}
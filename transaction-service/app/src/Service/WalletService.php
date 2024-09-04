<?php

namespace App\Service;

use App\Entity\Wallet;

class WalletService
{
    public static function syncWalletBalance(Wallet $wallet): void
    {
        $sum = 0;
        foreach ($wallet->getTransactions() as $transaction) {
            if ($transaction->getType() === 'income') {
                $sum += $transaction->getAmount();
            }
            if ($transaction->getType() === 'outcome') {
                $sum -= $transaction->getAmount();
            }
        }

        $wallet->setBalance($sum);
    }
}
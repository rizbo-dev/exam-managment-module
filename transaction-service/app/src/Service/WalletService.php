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
                $sum += 0;
            }
            if ($transaction->getType() === 'outcome') {
                $sum -= 0;
            }
        }

        $wallet->setBalance($sum);
    }
}
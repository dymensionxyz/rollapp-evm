'use client'

import { useState } from 'react';
import { motion } from 'framer-motion';
import { Button, TextField, Radio, RadioGroup, FormControlLabel } from '@mui/material';
import { toast } from 'react-hot-toast';

export default function CoinFlipGame() {
  const [bet, setBet] = useState(1);
  const [side, setSide] = useState('heads');
  const [balance, setBalance] = useState(100);
  const [isFlipping, setIsFlipping] = useState(false);
  const [result, setResult] = useState<'heads' | 'tails' | null>(null);
  const [error, setError] = useState('');

  const flipCoin = () => {
    if (bet > balance) {
      setError('Insufficient balance for this bet');
      return;
    }

    setError('');
    setIsFlipping(true);
    setResult(null);

    setTimeout(() => {
      const flipResult = Math.random() < 0.5 ? 'heads' : 'tails';
      setResult(flipResult);
      setIsFlipping(false);

      if (flipResult === side) {
        setBalance(balance + bet);
        toast.success(`You won $${bet}. Your new balance is $${balance + bet}.`);
      } else {
        setBalance(balance - bet);
        toast.error(`You lost $${bet}. Your new balance is $${balance - bet}.`);
      }
    }, 2000);
  };

  return (
      <div className="flex flex-col items-center justify-center min-h-screen bg-gradient-to-r from-purple-400 via-pink-500 to-red-500">
        <div className="p-8 bg-white rounded-lg shadow-xl w-full max-w-md">
          <h1 className="text-3xl font-bold text-center mb-6">Coin Flip Game</h1>

          {error && <div className="text-red-500 text-center mb-4">{error}</div>}

          <div className="mb-6 text-center">
            <motion.div
                className="w-32 h-32 rounded-full bg-yellow-400 mx-auto flex items-center justify-center"
                animate={{
                  rotateY: isFlipping ? 1800 : 0,
                  scale: isFlipping ? 1.2 : 1,
                }}
                transition={{ duration: 2 }}
            >
              {result && <div className="text-2xl font-bold">{result === 'heads' ? 'Heads' : 'Tails'}</div>}
            </motion.div>
          </div>

          <div className="space-y-4">
            <div>
              <TextField
                  label="Your Bet ($)"
                  type="number"
                  value={bet}
                  onChange={(e) => setBet(Math.max(1, parseInt(e.target.value)))}
                  fullWidth
                  inputProps={{ min: 1, max: balance }}
              />
            </div>

            <RadioGroup value={side} onChange={(e) => setSide(e.target.value)}>
              <FormControlLabel value="heads" control={<Radio />} label="Heads" />
              <FormControlLabel value="tails" control={<Radio />} label="Tails" />
            </RadioGroup>

            <Button variant="contained" onClick={flipCoin} disabled={isFlipping} fullWidth>
              {isFlipping ? 'Flipping...' : 'Flip Coin'}
            </Button>

            <div className="text-center text-xl font-semibold">Balance: ${balance}</div>
          </div>
        </div>
      </div>
  );
}

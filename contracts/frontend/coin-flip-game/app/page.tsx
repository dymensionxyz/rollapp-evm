'use client'

import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';
import { Button, TextField, Radio, RadioGroup, FormControlLabel } from '@mui/material';
import { toast } from 'react-hot-toast';
import { ethers } from 'ethers'; // Используем ethers для контрактов
import CoinFlipABI from './CoinFlipABI.json'; // Путь к ABI вашего контракта

function sleep(ms: number): Promise<void> {
  return new Promise(resolve => setTimeout(resolve, ms));
}

export default function CoinFlipGame() {
  const [bet, setBet] = useState(1);
  const [side, setSide] = useState('heads');
  const [balance, setBalance] = useState(100);
  const [isFlipping, setIsFlipping] = useState(false);
  const [result, setResult] = useState<'heads' | 'tails' | null>(null);
  const [error, setError] = useState('');
  const [provider, setProvider] = useState<any>(null); // Поставим any, чтобы не было проблем с типами
  const [signer, setSigner] = useState<any>(null);
  const [coinFlipContract, setCoinFlipContract] = useState<any>(null);
  const [connected, setConnected] = useState(false);

  const CONTRACT_ADDRESS = '0x5adfA443a4D6F70A0226dF8ADBD72aA78E14E256'; // Введите адрес задеплоенного контракта

  useEffect(() => {
    if (window.ethereum) {
      const initWeb3 = async () => {
        const ethereum = window.ethereum;
        // Проверка на доступность MetaMask
        if (typeof ethereum !== 'undefined') {
          try {
            const web3Provider = new ethers.BrowserProvider(ethereum);

            const network = await web3Provider.getNetwork();
            console.log('Connected to network:', network.name, 'with chainId', network.chainId);

            const userSigner = await web3Provider.getSigner();
            const contract = new ethers.Contract(CONTRACT_ADDRESS, CoinFlipABI, userSigner);

            setProvider(web3Provider);
            setSigner(userSigner);
            setCoinFlipContract(contract);
            setConnected(true);
          } catch (err) {
            console.error('Failed to connect to MetaMask', err);
            setError('Failed to connect to MetaMask');
          }
        } else {
          setError('Please install MetaMask');
        }
      };

      initWeb3()
    } else {
      setError('Please install MetaMask');
    }
  }, []);

  const flipCoin = async () => {
    console.log("flipping the coin with user choice: " + side)

    if (!connected) {
      setError('Please connect to MetaMask');
      return;
    }

    if (bet > balance) {
      setError('Insufficient balance for this bet');
      return;
    }

    setError('');
    setIsFlipping(true);
    setResult(null);

    try {
      const currentNonce = await provider.getTransactionCount(signer.address, "latest");
      const sideEnum = side === 'heads' ? 0 : 1; // 0 - HEADS, 1 - TAILS
      const tx = await coinFlipContract.startGame(sideEnum);
      await tx.wait()
      console.log('Game started');

      await completeGame()

    } catch (err) {
      console.error('Error flipping the coin:', err);
      setIsFlipping(false);
      setError('Error interacting with the contract');
    }
  };

  const completeGame = async () => {
    console.log(`completing the game`)
    try {
      const currentNonce = await provider.getTransactionCount(signer.address, "latest");

      var retry = 0
      while (retry < 3) {
        await sleep(1250);
        try {
          const tx = await coinFlipContract.completeGame({
            nonce: currentNonce
          });
          await tx.wait()
          break
        } catch (err) {
          retry++
          setError(`Error completing the game, retry #${retry}.`);
          console.error(`Retry reason: ${err}`)
        }
      }

      if (retry >= 3) {
        return;
      }

      const gameResult = await coinFlipContract.getPlayerLastGameResult();
      console.log(gameResult);

      if (gameResult.status != 2) {
        setError("Game wasn't finished.")
        console.error("Game wasn't finished.")
        return
      }
      const didWin = gameResult.won;
      const playerChoice = gameResult.playerChoice ? 'tails': 'heads'
      const flipResult = didWin ? playerChoice : (playerChoice === 'heads' ? 'tails' : 'heads');

      console.log("playerChoice: " + playerChoice + ", flipResult: " + flipResult)

      setResult(flipResult);
      setIsFlipping(false);

      if (didWin) {
        setBalance(balance + bet);
        toast.success(`You won $${bet}. Your new balance is $${balance + bet}.`);
      } else {
        setBalance(balance - bet);
        toast.error(`You lost $${bet}. Your new balance is $${balance - bet}.`);
      }
    } catch (err) {
      console.error('Error completing game:', err);
      setIsFlipping(false);
      setError('Error completing the game');
    }
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
                  rotateY: isFlipping ? 3600 : 0, // Бесконечное вращение, пока идет игра
                  scale: isFlipping ? 1.2 : 1,
                }}
                transition={{
                  duration: isFlipping ? 2 : 0, // Плавное вращение, если монета в процессе игры
                  repeat: isFlipping ? Infinity : 0, // Повторяет анимацию бесконечно
                  ease: 'easeInOut', // Плавный поворот
                }}
            >
              {result && <div className="text-2xl font-bold">{result === 'heads' ? 'heads' : 'tails'}</div>}
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
                  inputProps={{min: 1, max: balance}}
              />
            </div>

            <RadioGroup value={side} onChange={(e) => setSide(e.target.value)}>
              <FormControlLabel value="heads" control={<Radio/>} label="heads"/>
              <FormControlLabel value="tails" control={<Radio/>} label="tails"/>
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

'use client'

import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';
import { Button, TextField, Radio, RadioGroup, FormControlLabel } from '@mui/material';
import { toast } from 'react-hot-toast';
import { ethers } from 'ethers';
import CoinFlipABI from './CoinFlipABI.json';

function sleep(ms: number): Promise<void> {
  return new Promise(resolve => setTimeout(resolve, ms));
}

export default function CoinFlipGame() {
  const [bet, setBet] = useState(1);
  const [side, setSide] = useState('DYM');
  const [balance, setBalance] = useState(100);
  const [isFlipping, setIsFlipping] = useState(false);
  const [result, setResult] = useState<'DYM' | 'LOGO' | null>(null);
  const [error, setError] = useState('');
  const [provider, setProvider] = useState<any>(null);
  const [signer, setSigner] = useState<any>(null);
  const [coinFlipContract, setCoinFlipContract] = useState<any>(null);
  const [connected, setConnected] = useState(false);
  const [gameStatus, setGameStatus] = useState<'pending' | 'completed' | null>(null);
  const [winStatus, setWinStatus] = useState<'win' | 'lose' | null>(null)

  const CONTRACT_ADDRESS = '0x09d0647B434e6315f20AB0D6Cc87E1A274299b69';

  useEffect(() => {
    if (window.ethereum) {
      const initWeb3 = async () => {
        const ethereum = window.ethereum;
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
            await fetchGameStatus(contract);
          } catch (err) {
            console.error('Failed to connect to MetaMask', err);
            setError('Failed to connect to MetaMask');
          }
        } else {
          setError('Please install MetaMask');
        }
      };
      initWeb3();
    } else {
      setError('Please install MetaMask');
    }
  }, []);

  const fetchGameStatus = async (contract: any) => {
    try {
      const gameResult = await contract.getPlayerLastGameResult();
      console.log(gameResult);

      const status = gameResult.status.toString();
      let statusMessage = '';

      if (status === '2') {
        statusMessage = 'Completed';
        const didWin = gameResult.won;
        const playerChoice = gameResult.playerChoice ? 'LOGO' : 'DYM';
        const flipResult = didWin ? playerChoice : (playerChoice === 'DYM' ? 'LOGO' : 'DYM');
        console.log("playerChoice: " + playerChoice + ", flipResult: " + flipResult);

        if (didWin) {
          setWinStatus('win');
        } else {
          setWinStatus('lose');
        }

        setResult(flipResult);
        setIsFlipping(false);
        setGameStatus('completed');
      } else if (status === '1') {
        statusMessage = 'Pending';
      } else {
        statusMessage = 'No Game Started';
      }

      console.log('Fetched game status:', statusMessage);

      setGameStatus(status === '2' ? 'completed' : status === '1' ? 'pending' : null);
    } catch (err) {
      console.error('Error fetching game status:', err);
      setError('Error fetching game status');
    }
  };

  const flipCoin = async () => {
    console.log("flipping the coin with user choice: " + side);

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
    setGameStatus('pending');

    try {
      const currentNonce = await provider.getTransactionCount(signer.address, "latest");
      const sideEnum = side === 'DYM' ? 0 : 1;
      const tx = await coinFlipContract.startGame(sideEnum, {
        nonce: currentNonce,
        value: bet
      });
      await tx.wait();
      console.log('Game started');
      await completeGame(true);
    } catch (err) {
      console.error('Error flipping the coin:', err);
      setIsFlipping(false);
      setError('Error interacting with the contract');
      setGameStatus(null);
    }
  };

  const completeGame = async (wait: boolean) => {
    console.log(`completing the game`);
    try {
      var retry = 0;
      while (retry < 3) {
        if (wait) {
          await sleep(5000);
        }
        try {
          const currentNonce = await provider.getTransactionCount(signer.address, "latest");
          const tx = await coinFlipContract.completeGame({
            nonce: currentNonce
          });
          await tx.wait();
          break;
        } catch (err) {
          retry++;
          setError(`Error completing the game, retry #${retry}.`);
          console.error(`Retry reason: ${err}`);
        }
      }

      if (retry >= 3) {
        return;
      }

      const gameResult = await coinFlipContract.getPlayerLastGameResult();
      console.log(gameResult);

      if (gameResult.status != 2) {
        setError("Game wasn't finished.");
        console.error("Game wasn't finished.");
        return;
      }
      const didWin = gameResult.won;
      const playerChoice = gameResult.playerChoice ? 'LOGO' : 'DYM';
      const flipResult = didWin ? playerChoice : (playerChoice === 'DYM' ? 'LOGO' : 'DYM');

      console.log("playerChoice: " + playerChoice + ", flipResult: " + flipResult);

      setResult(flipResult);
      setIsFlipping(false);
      setGameStatus('completed');
      if (didWin) {
        setWinStatus('win');
      } else {
        setWinStatus('lose');
      }

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
      setGameStatus(null);
    }
  };

  return (
      <div className="flex flex-col items-center justify-center min-h-screen bg-gradient-to-r from-purple-400 via-pink-500 to-red-500">
        <div className="p-8 bg-white rounded-lg shadow-xl w-full max-w-md">
          <h1 className="text-3xl font-bold text-center mb-6">Coin Flip Game</h1>

          {error && <div className="text-red-500 text-center mb-4">{error}</div>}

          <div className="p-4 mb-4 bg-gray-100 rounded-lg shadow-md text-center">
            <h3 className="text-lg font-semibold">Game Status:</h3>
            <div className={`text-xl font-bold mt-2 ${gameStatus === 'pending' ? 'text-yellow-500' : gameStatus === 'completed' ? 'text-green-500' : 'text-gray-500'}`}>
              {gameStatus === 'pending' ? 'Pending...' : gameStatus === 'completed' ? 'Completed' : 'No Game Started'}
            </div>
            {gameStatus === 'completed' && (
                <div className={`text-xl font-bold mt-2 ${winStatus === 'win' ? 'text-green-500' : 'text-red-500'}`}>
                  {winStatus === 'win' ? 'You won!' : 'You lost!'}
                </div>
            )}
          </div>

          <div className="mb-6 text-center relative">
            <motion.div
                className="w-32 h-32 rounded-full mx-auto flex items-center justify-center relative"
                animate={{
                  rotateY: isFlipping ? 3600 : 0,
                  scale: isFlipping ? 1.2 : 1,
                }}
                transition={{
                  duration: isFlipping ? 2 : 0,
                  repeat: isFlipping ? Infinity : 0,
                  ease: 'easeInOut',
                }}
            >
              {/* Use an image of the coin */}
              <img
                  src="/coin.png"  // Ensure the path is correct to where your image is stored
                  alt="Coin"
                  className="w-full h-full object-contain"
              />
              {/* Text overlay */}
              {result && (
                  <div className="absolute inset-0 flex items-center justify-center text-2xl font-bold text-white">
                    {result === 'DYM' ? 'DYM' : 'LOGO'}
                  </div>
              )}
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
              <FormControlLabel value="DYM" control={<Radio />} label="DYM" />
              <FormControlLabel value="LOGO" control={<Radio />} label="LOGO" />
            </RadioGroup>

            <Button
                variant="contained"
                onClick={flipCoin}
                disabled={isFlipping || gameStatus === 'pending'}
                fullWidth
            >
              Flip
            </Button>

            <Button
                variant="contained"
                onClick={() => completeGame(false)}
                disabled={isFlipping || gameStatus !== 'pending'}
                fullWidth
            >
              Reveal
            </Button>

            <div className="text-center text-xl font-semibold">Balance: ${balance}</div>
          </div>
        </div>
      </div>
  );
}

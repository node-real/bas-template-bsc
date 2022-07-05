package eth

// TODO add abi
const gasFreeToAddressABI = `
[
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "constructorParams",
          "type": "bytes"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "prevValue",
          "type": "uint32"
        },
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "ActiveValidatorsLengthChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "prevValue",
          "type": "uint32"
        },
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "BurnRatioChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "prevValue",
          "type": "uint32"
        },
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "EpochBlockIntervalChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "prevValue",
          "type": "uint32"
        },
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "FelonyThresholdChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "freeGasToAddress",
          "type": "address"
        }
      ],
      "name": "FreeGasToAddressAdded",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "freeGasToAddress",
          "type": "address"
        }
      ],
      "name": "FreeGasToAddressRemoved",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "prevValue",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "newValue",
          "type": "uint256"
        }
      ],
      "name": "MinStakingAmountChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "prevValue",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "newValue",
          "type": "uint256"
        }
      ],
      "name": "MinValidatorStakeAmountChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "prevValue",
          "type": "uint32"
        },
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "MisdemeanorThresholdChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "prevValue",
          "type": "uint32"
        },
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "UndelegatePeriodChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "prevValue",
          "type": "uint32"
        },
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "ValidatorJailEpochLengthChanged",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "BURN_ADDRESS",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "BURN_RATIO_SCALE",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getChainConfig",
      "outputs": [
        {
          "internalType": "contract IChainConfig",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getGovernance",
      "outputs": [
        {
          "internalType": "contract IGovernance",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getSlashingIndicator",
      "outputs": [
        {
          "internalType": "contract ISlashingIndicator",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getStaking",
      "outputs": [
        {
          "internalType": "contract IStaking",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getStakingPool",
      "outputs": [
        {
          "internalType": "contract IStakingPool",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getSystemReward",
      "outputs": [
        {
          "internalType": "contract ISystemReward",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "init",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "contract IStaking",
          "name": "stakingContract",
          "type": "address"
        },
        {
          "internalType": "contract ISlashingIndicator",
          "name": "slashingIndicatorContract",
          "type": "address"
        },
        {
          "internalType": "contract ISystemReward",
          "name": "systemRewardContract",
          "type": "address"
        },
        {
          "internalType": "contract IStakingPool",
          "name": "stakingPoolContract",
          "type": "address"
        },
        {
          "internalType": "contract IGovernance",
          "name": "governanceContract",
          "type": "address"
        },
        {
          "internalType": "contract IChainConfig",
          "name": "chainConfigContract",
          "type": "address"
        },
        {
          "internalType": "contract IRuntimeUpgrade",
          "name": "runtimeUpgradeContract",
          "type": "address"
        },
        {
          "internalType": "contract IDeployerProxy",
          "name": "deployerProxyContract",
          "type": "address"
        }
      ],
      "name": "initManually",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "isInitialized",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint32",
          "name": "activeValidatorsLength",
          "type": "uint32"
        },
        {
          "internalType": "uint32",
          "name": "epochBlockInterval",
          "type": "uint32"
        },
        {
          "internalType": "uint32",
          "name": "misdemeanorThreshold",
          "type": "uint32"
        },
        {
          "internalType": "uint32",
          "name": "felonyThreshold",
          "type": "uint32"
        },
        {
          "internalType": "uint32",
          "name": "validatorJailEpochLength",
          "type": "uint32"
        },
        {
          "internalType": "uint32",
          "name": "undelegatePeriod",
          "type": "uint32"
        },
        {
          "internalType": "uint32",
          "name": "burnRatio",
          "type": "uint32"
        },
        {
          "internalType": "uint256",
          "name": "minValidatorStakeAmount",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "minStakingAmount",
          "type": "uint256"
        }
      ],
      "name": "ctor",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getActiveValidatorsLength",
      "outputs": [
        {
          "internalType": "uint32",
          "name": "",
          "type": "uint32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "setActiveValidatorsLength",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getEpochBlockInterval",
      "outputs": [
        {
          "internalType": "uint32",
          "name": "",
          "type": "uint32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "setEpochBlockInterval",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getMisdemeanorThreshold",
      "outputs": [
        {
          "internalType": "uint32",
          "name": "",
          "type": "uint32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "setMisdemeanorThreshold",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getFelonyThreshold",
      "outputs": [
        {
          "internalType": "uint32",
          "name": "",
          "type": "uint32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "setFelonyThreshold",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getValidatorJailEpochLength",
      "outputs": [
        {
          "internalType": "uint32",
          "name": "",
          "type": "uint32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "setValidatorJailEpochLength",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getUndelegatePeriod",
      "outputs": [
        {
          "internalType": "uint32",
          "name": "",
          "type": "uint32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "setUndelegatePeriod",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getBurnRatio",
      "outputs": [
        {
          "internalType": "uint32",
          "name": "",
          "type": "uint32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint32",
          "name": "newValue",
          "type": "uint32"
        }
      ],
      "name": "setBurnRatio",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getBurnRatioScale",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getBurnAddress",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getMinValidatorStakeAmount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "newValue",
          "type": "uint256"
        }
      ],
      "name": "setMinValidatorStakeAmount",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getMinStakingAmount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "newValue",
          "type": "uint256"
        }
      ],
      "name": "setMinStakingAmount",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "freeGasToAddress",
          "type": "address"
        }
      ],
      "name": "addFreeGasToAddress",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "freeGasToAddress",
          "type": "address"
        }
      ],
      "name": "removeFreeGasToAddress",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getFreeGasToAddressList",
      "outputs": [
        {
          "internalType": "address[]",
          "name": "",
          "type": "address[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }
  ]
`

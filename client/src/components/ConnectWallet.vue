<template>
    <div>
        <h1>Connect Wallet</h1>
        <p>Connect your wallet to continue</p>
        <button v-if="!isConnectedToMetaMask" @click="connectWeb3">Connect Wallet</button>
        <div v-else>
          <p>{{ userAddress }}</p>
          <p>{{ chainId }}</p>
        </div>
        <div>
          <p>New to our service? <button @click="registerSrv">Register</button></p>

          <p>Already registered?
            <button @click="loginSrv"> Login here</button>
          </p>
        </div>
    </div>
</template>
<script>

import {ref} from 'vue'
import {RequestParams} from '../config/request'
import axios from 'axios'

export default {
  name: 'ConnectWallet',
  setup () {
    const chainId = ref(0)
    return {
      chainId
    }
  },
  mounted () {
    if (localStorage.access_token) {
      this.access_token = localStorage.access_token
    }
  },
  watch: {
    access_token (val) {
      localStorage.access_token = val
    }
  },
  data () {
    return {
      userAddress: '',
      isConnectedToMetaMask: false,
      access_token: ''
    }
  },
  init () {
    this.connectWeb3()
  },
  methods: {
    connectWeb3: async function () {
      // Check if MetaMask is installed
      if (window.ethereum) {
        window.ethereum.request({ method: 'eth_requestAccounts' })
          .then(async (res) => {
            this.userAddress = res[0]
            await this.getUserAccount()
            this.isConnectedToMetaMask = true
          })
          .catch((err) => {
            console.error(err)
          })
      } else {
        // MetaMask is not installed, handle accordingly
        console.error('MetaMask is not installed.')
      }
    },
    getUserAccount: async function () {
      const res = await window.ethereum.request({
        method: 'eth_chainId'
      }).catch(err => {
        console.error(err)
      })
      this.chainId = parseInt(res, 16)
    },
    registerSrv () {
      this.registerServer({
        address: this.userAddress,
        chain_id: this.chainId.toString()
      })
    },
    loginSrv () {
      const data = {
        address: this.userAddress
      }
      axios.post(RequestParams.host + RequestParams.path.login, data, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(res => {
          this.access_token = res.data.access_token
          console.log(res)
        })
        .catch(err => {
          console.error(err)
        })
    },
    registerServer (data) {
      console.log('data sent:', data)
      axios.post(RequestParams.host + RequestParams.path.register, data, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(res => {
          console.log(res)
        })
        .catch(err => {
          console.error(err)
        })
    }
  }
}
</script>

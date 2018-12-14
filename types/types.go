package types

import (
	"github.com/tendermint/tendermint/libs/flowrate"
	"github.com/tendermint/tendermint/p2p/conn"
	"time"
)

type (
	BlockInfo struct {
		ID       string
		Height   int64
		Proposer string
		Time     time.Time
	}

	EvidenceInfo struct {
		Address string
		Height  int64
	}

	MissInfo struct {
		ID       int64
		Address  string
		Height   int64
		Alerted  bool `sql:",default:false,notnull"`
		Proposer string
		Time     time.Time
	}

	Proposal struct {
		Type    string `json:"type"`
		Alerted bool   `sql:",default:false,notnull"`
		Value   struct {
			ProposalID     string `json:"proposal_id"`
			Title          string `json:"title"`
			Description    string `json:"description"`
			ProposalType   string `json:"proposal_type"`
			ProposalStatus string `json:"proposal_status"`
			TallyResult    struct {
				Yes        string `json:"yes"`
				Abstain    string `json:"abstain"`
				No         string `json:"no"`
				NoWithVeto string `json:"no_with_veto"`
			} `json:"tally_result"`
			SubmitTime     string `json:"submit_time"`
			DepositEndTime string `json:"deposit_end_time"`
			TotalDeposit   struct {
				Denom  string `json:"denom"`
				Amount string `json:"amount"`
			} `json:"total_deposit"`
			VotingStartTime string `json:"voting_start_time"`
			VotingEndTime   string `json:"voting_end_time"`
		} `json:"value"`
	}

	PeerInfo struct {
		ID        int64
		Timestamp time.Time
		Node      string

		PeerID     string `json:"id"`
		ListenAddr string `json:"listen_addr"`
		Network    string `json:"network"`
		Version    string `json:"version"`
		Channels   string `json:"channels"`
		Moniker    string `json:"moniker"`
		IsOutbound bool   `json:"is_outbound";sql:",default:false,notnull"`

		SendData    flowrate.Status
		RecvData    flowrate.Status
		ChannelData []conn.ChannelStatus
	}
)

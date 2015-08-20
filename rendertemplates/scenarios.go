package rendertemplates

// TestScenarioData returns some example data
func TestScenarioData() *Tiers {
	return &Tiers{
		&Tier{
			Name: "try-anything",
			Slots: Slots{
				&Slot{
					Deployments: Deployments{
						&Deployment{
							Name: "try-anything / bosh-lite - cf-try-anything",
							Releases: []DisplayNameVersion{
								DisplayNameVersion{Name: "cf", Version: "214", DisplayClass: "icon-arrow-up green"},
								DisplayNameVersion{Name: "cf-haproxy", Version: "5", DisplayClass: "icon-minus blue"},
							},
							Stemcells: []DisplayNameVersion{
								DisplayNameVersion{Name: "warden", Version: "2776", DisplayClass: "icon-minus blue"},
							},
						},
						&Deployment{
							Name: "try-anything / bosh-lite - cf2-try-anything",
							Releases: []DisplayNameVersion{
								DisplayNameVersion{Name: "cf", Version: "215", DisplayClass: "icon-arrow-up green"},
								DisplayNameVersion{Name: "cf-haproxy", Version: "6", DisplayClass: "icon-arrow-up green"},
							},
							Stemcells: []DisplayNameVersion{
								DisplayNameVersion{Name: "warden", Version: "2776", DisplayClass: "icon-minus blue"},
							},
						},
					},
				},
			},
		},
		&Tier{
			Name: "dc",
			Slots: Slots{
				&Slot{
					Deployments: Deployments{
						&Deployment{
							Name: "dc / sandbox - cf-vsph-sandbox",
							Releases: []DisplayNameVersion{
								DisplayNameVersion{Name: "cf", Version: "215", DisplayClass: "icon-arrow-down red"},
								DisplayNameVersion{Name: "cf-haproxy", Version: "5", DisplayClass: "icon-minus blue"},
							},
							Stemcells: []DisplayNameVersion{
								DisplayNameVersion{Name: "vsphere", Version: "3048", DisplayClass: "icon-arrow-up green"},
							},
						},
					},
				},
				&Slot{},
				&Slot{},
			},
		},
		&Tier{
			Name: "aws",
			Slots: Slots{
				&Slot{},
				&Slot{},
				&Slot{
					Deployments: Deployments{
						&Deployment{
							Name: "aws / production - cf-aws-prod",
							Releases: []DisplayNameVersion{
								DisplayNameVersion{Name: "cf", Version: "211", DisplayClass: "icon-arrow-down red"},
								DisplayNameVersion{Name: "cf-haproxy", Version: "5", DisplayClass: "icon-minus blue"},
							},
							Stemcells: []DisplayNameVersion{
								DisplayNameVersion{Name: "aws", Version: "3033", DisplayClass: "icon-minus blue"},
							},
						},
					},
				},
			},
		},
	}
}

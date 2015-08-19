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
							Releases: []NameVersion{
								NameVersion{Name: "cf", Version: "214", DisplayClass: "icon-arrow-up green"},
								NameVersion{Name: "cf-haproxy", Version: "5", DisplayClass: "icon-minus blue"},
							},
							Stemcells: []NameVersion{
								NameVersion{Name: "warden", Version: "2776", DisplayClass: "icon-minus blue"},
							},
						},
						&Deployment{
							Name: "try-anything / bosh-lite - cf2-try-anything",
							Releases: []NameVersion{
								NameVersion{Name: "cf", Version: "215", DisplayClass: "icon-arrow-up green"},
								NameVersion{Name: "cf-haproxy", Version: "6", DisplayClass: "icon-arrow-up green"},
							},
							Stemcells: []NameVersion{
								NameVersion{Name: "warden", Version: "2776", DisplayClass: "icon-minus blue"},
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
							Name: "dc / sandbox / vsphere - cf-vsph-sandbox",
							Releases: []NameVersion{
								NameVersion{Name: "cf", Version: "215", DisplayClass: "icon-arrow-down red"},
								NameVersion{Name: "cf-haproxy", Version: "5", DisplayClass: "icon-minus blue"},
							},
							Stemcells: []NameVersion{
								NameVersion{Name: "vsphere", Version: "3048", DisplayClass: "icon-arrow-up green"},
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
							Name: "aws / sandbox / aws - cf-aws-prod",
							Releases: []NameVersion{
								NameVersion{Name: "cf", Version: "211", DisplayClass: "icon-arrow-down red"},
								NameVersion{Name: "cf-haproxy", Version: "5", DisplayClass: "icon-minus blue"},
							},
							Stemcells: []NameVersion{
								NameVersion{Name: "aws", Version: "3033", DisplayClass: "icon-minus blue"},
							},
						},
					},
				},
			},
		},
	}
}

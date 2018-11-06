package websites

import (
	"fmt"
	"testing"

	"github.com/tombuildsstuff/golang-iis/iis/applicationpools"
	"github.com/tombuildsstuff/golang-iis/iis/cmd"
	"github.com/tombuildsstuff/golang-iis/iis/helpers"
)

func TestWebsiteLifecycle(t *testing.T) {
	rInt := helpers.RandomInt()
	appPoolName := fmt.Sprintf("acctestpool-%d", rInt)
	websiteName := fmt.Sprintf("acctestsite-%d", rInt)

	appPoolsClient := applicationpools.AppPoolsClient{
		Client: cmd.Client{},
	}
	websitesClient := WebsitesClient{
		Client: cmd.Client{},
	}

	err := appPoolsClient.Create(appPoolName)
	if err != nil {
		t.Fatalf("Error creating App Pool %q: %+v", appPoolName, err)
		return
	}

	err = websitesClient.Create(websiteName, appPoolName, defaultWebsitePath)
	if err != nil {
		t.Fatalf("Error creating Website %q in App Pool %q: %+v", websiteName, appPoolName, err)
		return
	}

	site, err := websitesClient.Get(websiteName)
	if err != nil {
		t.Fatalf("Error retrieving Website %q (App Pool %q): %+v", websiteName, appPoolName, err)
		return
	}

	if site.Name != websiteName {
		t.Fatalf("Expected the Name to be %q but got %q", websiteName, site.Name)
		return
	}

	if site.PhysicalPath != defaultWebsitePath {
		t.Fatalf("Expected the Physical Path to be %q but got %q", defaultWebsitePath, site.PhysicalPath)
		return
	}

	if site.ApplicationPool != appPoolName {
		t.Fatalf("Expected the App Pool to be %q but got %q", appPoolName, site.ApplicationPool)
		return
	}

	if site.StartsOnBoot != false {
		t.Fatalf("Expedted StartOnBoot to be false but it wasn't!")
		return
	}

	// TODO: alter the websites state and confirm it
}

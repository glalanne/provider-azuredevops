// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	agentpool "github.com/glalanne/provider-azuredevops/internal/controller/cluster/agents/agentpool"
	agentqueue "github.com/glalanne/provider-azuredevops/internal/controller/cluster/agents/agentqueue"
	elasticpool "github.com/glalanne/provider-azuredevops/internal/controller/cluster/agents/elasticpool"
	feed "github.com/glalanne/provider-azuredevops/internal/controller/cluster/artifacts/feed"
	feedretentionpolicy "github.com/glalanne/provider-azuredevops/internal/controller/cluster/artifacts/feedretentionpolicy"
	dashboard "github.com/glalanne/provider-azuredevops/internal/controller/cluster/boards/dashboard"
	extension "github.com/glalanne/provider-azuredevops/internal/controller/cluster/core/extension"
	builddefinition "github.com/glalanne/provider-azuredevops/internal/controller/cluster/pipelines/builddefinition"
	buildfolder "github.com/glalanne/provider-azuredevops/internal/controller/cluster/pipelines/buildfolder"
	environment "github.com/glalanne/provider-azuredevops/internal/controller/cluster/pipelines/environment"
	servicehookstoragequeuepipelines "github.com/glalanne/provider-azuredevops/internal/controller/cluster/pipelines/servicehookstoragequeuepipelines"
	variablegroup "github.com/glalanne/provider-azuredevops/internal/controller/cluster/pipelines/variablegroup"
	project "github.com/glalanne/provider-azuredevops/internal/controller/cluster/project/project"
	projectfeatures "github.com/glalanne/provider-azuredevops/internal/controller/cluster/project/projectfeatures"
	projectpipelinesettings "github.com/glalanne/provider-azuredevops/internal/controller/cluster/project/projectpipelinesettings"
	projecttags "github.com/glalanne/provider-azuredevops/internal/controller/cluster/project/projecttags"
	wiki "github.com/glalanne/provider-azuredevops/internal/controller/cluster/project/wiki"
	wikipage "github.com/glalanne/provider-azuredevops/internal/controller/cluster/project/wikipage"
	workitem "github.com/glalanne/provider-azuredevops/internal/controller/cluster/project/workitem"
	providerconfig "github.com/glalanne/provider-azuredevops/internal/controller/cluster/providerconfig"
	branchpolicyautoreviewers "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/branchpolicyautoreviewers"
	branchpolicybuildvalidation "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/branchpolicybuildvalidation"
	branchpolicycommentresolution "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/branchpolicycommentresolution"
	branchpolicymergetypes "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/branchpolicymergetypes"
	branchpolicyminreviewers "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/branchpolicyminreviewers"
	branchpolicystatuscheck "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/branchpolicystatuscheck"
	branchpolicyworkitemlinking "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/branchpolicyworkitemlinking"
	checkapproval "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/checkapproval"
	checkbranchcontrol "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/checkbranchcontrol"
	checkbusinesshours "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/checkbusinesshours"
	checkexclusivelock "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/checkexclusivelock"
	checkrequiredtemplate "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/checkrequiredtemplate"
	checkrestapi "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/checkrestapi"
	gitrepository "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/gitrepository"
	gitrepositorybranch "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/gitrepositorybranch"
	gitrepositoryfile "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/gitrepositoryfile"
	repositorypolicyauthoremailpattern "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/repositorypolicyauthoremailpattern"
	repositorypolicycaseenforcement "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/repositorypolicycaseenforcement"
	repositorypolicycheckcredentials "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/repositorypolicycheckcredentials"
	repositorypolicyfilepathpattern "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/repositorypolicyfilepathpattern"
	repositorypolicymaxfilesize "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/repositorypolicymaxfilesize"
	repositorypolicymaxpathlength "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/repositorypolicymaxpathlength"
	repositorypolicyreservednames "github.com/glalanne/provider-azuredevops/internal/controller/cluster/repos/repositorypolicyreservednames"
	areapermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/areapermissions"
	builddefinitionpermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/builddefinitionpermissions"
	buildfolderpermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/buildfolderpermissions"
	feedpermission "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/feedpermission"
	gitpermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/gitpermissions"
	group "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/group"
	groupentitlement "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/groupentitlement"
	groupmembership "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/groupmembership"
	iterationpermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/iterationpermissions"
	librarypermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/librarypermissions"
	pipelineauthorization "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/pipelineauthorization"
	projectpermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/projectpermissions"
	resourceauthorization "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/resourceauthorization"
	securityroleassignment "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/securityroleassignment"
	serviceendpointpermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/serviceendpointpermissions"
	servicehookpermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/servicehookpermissions"
	serviceprincipalentitlement "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/serviceprincipalentitlement"
	taggingpermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/taggingpermissions"
	team "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/team"
	teamadministrators "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/teamadministrators"
	teammembers "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/teammembers"
	userentitlement "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/userentitlement"
	variablegrouppermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/variablegrouppermissions"
	workitemquerypermissions "github.com/glalanne/provider-azuredevops/internal/controller/cluster/security/workitemquerypermissions"
	serviceendpointargocd "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointargocd"
	serviceendpointartifactory "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointartifactory"
	serviceendpointaws "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointaws"
	serviceendpointazurecr "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointazurecr"
	serviceendpointazuredevops "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointazuredevops"
	serviceendpointazurerm "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointazurerm"
	serviceendpointazureservicebus "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointazureservicebus"
	serviceendpointbitbucket "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointbitbucket"
	serviceendpointblackduck "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointblackduck"
	serviceendpointcheckmarxone "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointcheckmarxone"
	serviceendpointcheckmarxsast "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointcheckmarxsast"
	serviceendpointcheckmarxsca "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointcheckmarxsca"
	serviceendpointdockerregistry "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointdockerregistry"
	serviceendpointdynamicslifecycleservices "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointdynamicslifecycleservices"
	serviceendpointexternaltfs "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointexternaltfs"
	serviceendpointgcpterraform "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointgcpterraform"
	serviceendpointgeneric "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointgeneric"
	serviceendpointgenericgit "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointgenericgit"
	serviceendpointgithub "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointgithub"
	serviceendpointgithubenterprise "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointgithubenterprise"
	serviceendpointgitlab "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointgitlab"
	serviceendpointincomingwebhook "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointincomingwebhook"
	serviceendpointjenkins "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointjenkins"
	serviceendpointjfrogartifactoryv2 "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointjfrogartifactoryv2"
	serviceendpointjfrogdistributionv2 "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointjfrogdistributionv2"
	serviceendpointjfrogplatformv2 "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointjfrogplatformv2"
	serviceendpointjfrogxrayv2 "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointjfrogxrayv2"
	serviceendpointkubernetes "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointkubernetes"
	serviceendpointmaven "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointmaven"
	serviceendpointnexus "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointnexus"
	serviceendpointnpm "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointnpm"
	serviceendpointnuget "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointnuget"
	serviceendpointoctopusdeploy "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointoctopusdeploy"
	serviceendpointopenshift "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointopenshift"
	serviceendpointrunpipeline "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointrunpipeline"
	serviceendpointservicefabric "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointservicefabric"
	serviceendpointsnyk "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointsnyk"
	serviceendpointsonarcloud "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointsonarcloud"
	serviceendpointsonarqube "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointsonarqube"
	serviceendpointssh "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointssh"
	serviceendpointvisualstudiomarketplace "github.com/glalanne/provider-azuredevops/internal/controller/cluster/serviceendpoint/serviceendpointvisualstudiomarketplace"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agentpool.Setup,
		agentqueue.Setup,
		elasticpool.Setup,
		feed.Setup,
		feedretentionpolicy.Setup,
		dashboard.Setup,
		extension.Setup,
		builddefinition.Setup,
		buildfolder.Setup,
		environment.Setup,
		servicehookstoragequeuepipelines.Setup,
		variablegroup.Setup,
		project.Setup,
		projectfeatures.Setup,
		projectpipelinesettings.Setup,
		projecttags.Setup,
		wiki.Setup,
		wikipage.Setup,
		workitem.Setup,
		providerconfig.Setup,
		branchpolicyautoreviewers.Setup,
		branchpolicybuildvalidation.Setup,
		branchpolicycommentresolution.Setup,
		branchpolicymergetypes.Setup,
		branchpolicyminreviewers.Setup,
		branchpolicystatuscheck.Setup,
		branchpolicyworkitemlinking.Setup,
		checkapproval.Setup,
		checkbranchcontrol.Setup,
		checkbusinesshours.Setup,
		checkexclusivelock.Setup,
		checkrequiredtemplate.Setup,
		checkrestapi.Setup,
		gitrepository.Setup,
		gitrepositorybranch.Setup,
		gitrepositoryfile.Setup,
		repositorypolicyauthoremailpattern.Setup,
		repositorypolicycaseenforcement.Setup,
		repositorypolicycheckcredentials.Setup,
		repositorypolicyfilepathpattern.Setup,
		repositorypolicymaxfilesize.Setup,
		repositorypolicymaxpathlength.Setup,
		repositorypolicyreservednames.Setup,
		areapermissions.Setup,
		builddefinitionpermissions.Setup,
		buildfolderpermissions.Setup,
		feedpermission.Setup,
		gitpermissions.Setup,
		group.Setup,
		groupentitlement.Setup,
		groupmembership.Setup,
		iterationpermissions.Setup,
		librarypermissions.Setup,
		pipelineauthorization.Setup,
		projectpermissions.Setup,
		resourceauthorization.Setup,
		securityroleassignment.Setup,
		serviceendpointpermissions.Setup,
		servicehookpermissions.Setup,
		serviceprincipalentitlement.Setup,
		taggingpermissions.Setup,
		team.Setup,
		teamadministrators.Setup,
		teammembers.Setup,
		userentitlement.Setup,
		variablegrouppermissions.Setup,
		workitemquerypermissions.Setup,
		serviceendpointargocd.Setup,
		serviceendpointartifactory.Setup,
		serviceendpointaws.Setup,
		serviceendpointazurecr.Setup,
		serviceendpointazuredevops.Setup,
		serviceendpointazurerm.Setup,
		serviceendpointazureservicebus.Setup,
		serviceendpointbitbucket.Setup,
		serviceendpointblackduck.Setup,
		serviceendpointcheckmarxone.Setup,
		serviceendpointcheckmarxsast.Setup,
		serviceendpointcheckmarxsca.Setup,
		serviceendpointdockerregistry.Setup,
		serviceendpointdynamicslifecycleservices.Setup,
		serviceendpointexternaltfs.Setup,
		serviceendpointgcpterraform.Setup,
		serviceendpointgeneric.Setup,
		serviceendpointgenericgit.Setup,
		serviceendpointgithub.Setup,
		serviceendpointgithubenterprise.Setup,
		serviceendpointgitlab.Setup,
		serviceendpointincomingwebhook.Setup,
		serviceendpointjenkins.Setup,
		serviceendpointjfrogartifactoryv2.Setup,
		serviceendpointjfrogdistributionv2.Setup,
		serviceendpointjfrogplatformv2.Setup,
		serviceendpointjfrogxrayv2.Setup,
		serviceendpointkubernetes.Setup,
		serviceendpointmaven.Setup,
		serviceendpointnexus.Setup,
		serviceendpointnpm.Setup,
		serviceendpointnuget.Setup,
		serviceendpointoctopusdeploy.Setup,
		serviceendpointopenshift.Setup,
		serviceendpointrunpipeline.Setup,
		serviceendpointservicefabric.Setup,
		serviceendpointsnyk.Setup,
		serviceendpointsonarcloud.Setup,
		serviceendpointsonarqube.Setup,
		serviceendpointssh.Setup,
		serviceendpointvisualstudiomarketplace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agentpool.SetupGated,
		agentqueue.SetupGated,
		elasticpool.SetupGated,
		feed.SetupGated,
		feedretentionpolicy.SetupGated,
		dashboard.SetupGated,
		extension.SetupGated,
		builddefinition.SetupGated,
		buildfolder.SetupGated,
		environment.SetupGated,
		servicehookstoragequeuepipelines.SetupGated,
		variablegroup.SetupGated,
		project.SetupGated,
		projectfeatures.SetupGated,
		projectpipelinesettings.SetupGated,
		projecttags.SetupGated,
		wiki.SetupGated,
		wikipage.SetupGated,
		workitem.SetupGated,
		providerconfig.SetupGated,
		branchpolicyautoreviewers.SetupGated,
		branchpolicybuildvalidation.SetupGated,
		branchpolicycommentresolution.SetupGated,
		branchpolicymergetypes.SetupGated,
		branchpolicyminreviewers.SetupGated,
		branchpolicystatuscheck.SetupGated,
		branchpolicyworkitemlinking.SetupGated,
		checkapproval.SetupGated,
		checkbranchcontrol.SetupGated,
		checkbusinesshours.SetupGated,
		checkexclusivelock.SetupGated,
		checkrequiredtemplate.SetupGated,
		checkrestapi.SetupGated,
		gitrepository.SetupGated,
		gitrepositorybranch.SetupGated,
		gitrepositoryfile.SetupGated,
		repositorypolicyauthoremailpattern.SetupGated,
		repositorypolicycaseenforcement.SetupGated,
		repositorypolicycheckcredentials.SetupGated,
		repositorypolicyfilepathpattern.SetupGated,
		repositorypolicymaxfilesize.SetupGated,
		repositorypolicymaxpathlength.SetupGated,
		repositorypolicyreservednames.SetupGated,
		areapermissions.SetupGated,
		builddefinitionpermissions.SetupGated,
		buildfolderpermissions.SetupGated,
		feedpermission.SetupGated,
		gitpermissions.SetupGated,
		group.SetupGated,
		groupentitlement.SetupGated,
		groupmembership.SetupGated,
		iterationpermissions.SetupGated,
		librarypermissions.SetupGated,
		pipelineauthorization.SetupGated,
		projectpermissions.SetupGated,
		resourceauthorization.SetupGated,
		securityroleassignment.SetupGated,
		serviceendpointpermissions.SetupGated,
		servicehookpermissions.SetupGated,
		serviceprincipalentitlement.SetupGated,
		taggingpermissions.SetupGated,
		team.SetupGated,
		teamadministrators.SetupGated,
		teammembers.SetupGated,
		userentitlement.SetupGated,
		variablegrouppermissions.SetupGated,
		workitemquerypermissions.SetupGated,
		serviceendpointargocd.SetupGated,
		serviceendpointartifactory.SetupGated,
		serviceendpointaws.SetupGated,
		serviceendpointazurecr.SetupGated,
		serviceendpointazuredevops.SetupGated,
		serviceendpointazurerm.SetupGated,
		serviceendpointazureservicebus.SetupGated,
		serviceendpointbitbucket.SetupGated,
		serviceendpointblackduck.SetupGated,
		serviceendpointcheckmarxone.SetupGated,
		serviceendpointcheckmarxsast.SetupGated,
		serviceendpointcheckmarxsca.SetupGated,
		serviceendpointdockerregistry.SetupGated,
		serviceendpointdynamicslifecycleservices.SetupGated,
		serviceendpointexternaltfs.SetupGated,
		serviceendpointgcpterraform.SetupGated,
		serviceendpointgeneric.SetupGated,
		serviceendpointgenericgit.SetupGated,
		serviceendpointgithub.SetupGated,
		serviceendpointgithubenterprise.SetupGated,
		serviceendpointgitlab.SetupGated,
		serviceendpointincomingwebhook.SetupGated,
		serviceendpointjenkins.SetupGated,
		serviceendpointjfrogartifactoryv2.SetupGated,
		serviceendpointjfrogdistributionv2.SetupGated,
		serviceendpointjfrogplatformv2.SetupGated,
		serviceendpointjfrogxrayv2.SetupGated,
		serviceendpointkubernetes.SetupGated,
		serviceendpointmaven.SetupGated,
		serviceendpointnexus.SetupGated,
		serviceendpointnpm.SetupGated,
		serviceendpointnuget.SetupGated,
		serviceendpointoctopusdeploy.SetupGated,
		serviceendpointopenshift.SetupGated,
		serviceendpointrunpipeline.SetupGated,
		serviceendpointservicefabric.SetupGated,
		serviceendpointsnyk.SetupGated,
		serviceendpointsonarcloud.SetupGated,
		serviceendpointsonarqube.SetupGated,
		serviceendpointssh.SetupGated,
		serviceendpointvisualstudiomarketplace.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

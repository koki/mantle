package volume

import (
	"fmt"
	"reflect"

	. "mantle/pkg/core/pod/volume/aws"
	. "mantle/pkg/core/pod/volume/azure"
	. "mantle/pkg/core/pod/volume/ceph"
	. "mantle/pkg/core/pod/volume/cinder"
	. "mantle/pkg/core/pod/volume/configmap"
	. "mantle/pkg/core/pod/volume/downwardapi"
	. "mantle/pkg/core/pod/volume/emptydir"
	. "mantle/pkg/core/pod/volume/fibrechannel"
	. "mantle/pkg/core/pod/volume/flex"
	. "mantle/pkg/core/pod/volume/flocker"
	. "mantle/pkg/core/pod/volume/gcepd"
	. "mantle/pkg/core/pod/volume/git"
	. "mantle/pkg/core/pod/volume/gluster"
	. "mantle/pkg/core/pod/volume/hostpath"
	. "mantle/pkg/core/pod/volume/iscsi"
	. "mantle/pkg/core/pod/volume/nfs"
	. "mantle/pkg/core/pod/volume/photon"
	. "mantle/pkg/core/pod/volume/portworx"
	. "mantle/pkg/core/pod/volume/projected"
	. "mantle/pkg/core/pod/volume/pvc"
	. "mantle/pkg/core/pod/volume/quobyte"
	. "mantle/pkg/core/pod/volume/rbd"
	. "mantle/pkg/core/pod/volume/scaleio"
	. "mantle/pkg/core/pod/volume/secret"
	. "mantle/pkg/core/pod/volume/storageos"
	. "mantle/pkg/core/pod/volume/vsphere"

	"k8s.io/api/core/v1"
)

type Volume struct {
	HostPath     *HostPathVolume
	EmptyDir     *EmptyDirVolume
	GcePD        *GcePDVolume
	AwsEBS       *AwsEBSVolume
	AzureDisk    *AzureDiskVolume
	AzureFile    *AzureFileVolume
	CephFS       *CephFSVolume
	Cinder       *CinderVolume
	FibreChannel *FibreChannelVolume
	Flex         *FlexVolume
	Flocker      *FlockerVolume
	Glusterfs    *GlusterfsVolume
	ISCSI        *ISCSIVolume
	NFS          *NFSVolume
	PhotonPD     *PhotonPDVolume
	Portworx     *PortworxVolume
	PVC          *PVCVolume
	Quobyte      *QuobyteVolume
	ScaleIO      *ScaleIOVolume
	Vsphere      *VsphereVolume
	ConfigMap    *ConfigMapVolume
	Secret       *SecretVolume
	DownwardAPI  *DownwardAPIVolume
	Projected    *ProjectedVolume
	Git          *GitVolume
	RBD          *RBDVolume
	StorageOS    *StorageOSVolume
}

func (v *Volume) ToKube(version string) (interface{}, error) {
	fields := reflect.ValueOf(v).Elem()
	for n := 0; n < fields.NumField(); n++ {
		field := fields.Field(n)
		if field.IsValid() && !field.IsNil() {
			var err error

			convFunc := field.MethodByName("ToKube")
			resp := convFunc.Call([]reflect.Value{reflect.ValueOf(version)})
			if !resp[1].IsNil() {
				err = resp[1].Interface().(error)
			} else {
				err = nil
			}

			if !resp[0].IsNil() {
				return resp[0].Interface(), err
			}
			return nil, err
		}
	}

	return nil, fmt.Errorf("no volume type set")
}

func NewVolumeFromKubeVolume(obj interface{}) (*Volume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.Volume{}):
		o := obj.(v1.Volume)
		return fromKubeVolumeV1(&o)
	case reflect.TypeOf(&v1.Volume{}):
		return fromKubeVolumeV1(obj.(*v1.Volume))
	default:
		return nil, fmt.Errorf("unknown Volume version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeVolumeV1(obj *v1.Volume) (*Volume, error) {
	var vol Volume
	var err error

	if obj.HostPath != nil {
		vol.HostPath, err = NewHostPathVolumeFromKubeHostPathVolumeSource(obj.HostPath)
	}

	if obj.EmptyDir != nil {
		vol.EmptyDir, err = NewEmptyDirVolumeFromKubeEmptyDirVolumeSource(obj.EmptyDir)
	}

	if obj.GCEPersistentDisk != nil {
		vol.GcePD, err = NewGcePDVolumeFromKubeGCEPersistentDiskVolumeSource(obj.GCEPersistentDisk)
	}

	if obj.AWSElasticBlockStore != nil {
		vol.AwsEBS, err = NewAwsEBSVolumeFromKubeAWSElasticBlockStoreVolumeSource(obj.AWSElasticBlockStore)
	}

	if obj.GitRepo != nil {
		vol.Git, err = NewGitVolumeFromKubeGitRepoVolumeSource(obj.GitRepo)
	}

	if obj.Secret != nil {
		vol.Secret, err = NewSecretVolumeFromKubeSecretVolumeSource(obj.Secret)
	}

	if obj.NFS != nil {
		vol.NFS, err = NewNFSVolumeFromNFSVolumeSource(obj.NFS)
	}

	if obj.ISCSI != nil {
		vol.ISCSI, err = NewAwsEBSVolumeFromKubeISCSIVolumeSource(obj.ISCSI)
	}

	if obj.Glusterfs != nil {
		vol.Glusterfs, err = NewGlusterfsVolumeFromKubeGlusterfsVolumeSource(obj.Glusterfs)
	}

	if obj.PersistentVolumeClaim != nil {
		vol.PVC, err = NewPVCVolumeFromKubePersistentVolumeClaimVolumeSource(obj.PersistentVolumeClaim)
	}

	if obj.RBD != nil {
		vol.RBD, err = NewRBDVolumeFromKubeRBDVolumeSource(obj.RBD)
	}

	if obj.FlexVolume != nil {
		vol.Flex, err = NewFlexVolumeFromKubeFlexVolumeSource(obj.FlexVolume)
	}

	if obj.Cinder != nil {
		vol.Cinder, err = NewCinderVolumeFromKubeCinderVolumeSource(obj.Cinder)
	}

	if obj.CephFS != nil {
		vol.CephFS, err = NewCephFSVolumeFromKubeCephFSVolumeSource(obj.CephFS)
	}

	if obj.Flocker != nil {
		vol.Flocker, err = NewFlockerVolumeFromKubeFlockerVolumeSource(obj.Flocker)
	}

	if obj.DownwardAPI != nil {
		vol.DownwardAPI, err = NewDownwardAPIVolumeFromKubeDownwardAPIVolumeSource(obj.DownwardAPI)
	}

	if obj.FC != nil {
		vol.FibreChannel, err = NewFibreChannelVolumeFromKubeFCVolumeSource(obj.FC)
	}

	if obj.AzureFile != nil {
		vol.AzureFile, err = NewAzureFileVolumeFromAzureFileVolumeSource(obj.AzureFile)
	}

	if obj.ConfigMap != nil {
		vol.ConfigMap, err = NewConfigMapVolumeFromKubeConfigMapVolumeSource(obj.ConfigMap)
	}

	if obj.VsphereVolume != nil {
		vol.Vsphere, err = NewVsphereVolumeFromKubeVsphereVirtualDiskVolumeSource(obj.VsphereVolume)
	}

	if obj.Quobyte != nil {
		vol.Quobyte, err = NewQuobyteVolumeFromKubeQuobyteVolumeSource(obj.Quobyte)
	}

	if obj.AzureDisk != nil {
		vol.AzureDisk, err = NewAzureDiskVolumeFromAzureDiskVolumeSource(obj.AzureDisk)
	}

	if obj.PhotonPersistentDisk != nil {
		vol.PhotonPD, err = NewPhotonPDVolumeFromKubePhotonPersistentDiskVolumeSource(obj.PhotonPersistentDisk)
	}

	if obj.Projected != nil {
		vol.Projected, err = NewProjectedVolumeFromKubeProjectedVolumeSource(obj.Projected)
	}

	if obj.PortworxVolume != nil {
		vol.Portworx, err = NewPortworxVolumeVolumeFromKubePortworxVolumeSource(obj.PortworxVolume)
	}

	if obj.ScaleIO != nil {
		vol.ScaleIO, err = NewScaleIOVolumeFromKubeScaleIOVolumeSource(obj.ScaleIO)
	}

	if obj.StorageOS != nil {
		vol.StorageOS, err = NewStorageOSVolumeFromKubeStorageOSVolumeSource(obj.StorageOS)
	}

	return &vol, err
}

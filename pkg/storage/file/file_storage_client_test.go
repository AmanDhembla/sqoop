package file_test

import (
	"io/ioutil"
	"os"

	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/pkg/log"
	. "github.com/solo-io/gloo/pkg/storage/file"
	. "github.com/solo-io/gloo/test/helpers"
)

var _ = Describe("CrdStorageClient", func() {
	var (
		dir    string
		err    error
		resync = time.Second
	)
	BeforeEach(func() {
		dir, err = ioutil.TempDir("", "filecachetest")
		Must(err)
	})
	AfterEach(func() {
		log.Debugf("removing " + dir)
		os.RemoveAll(dir)
	})
	Describe("New", func() {
		It("creates a new client without error", func() {
			_, err = NewStorage(dir, resync)
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Create", func() {
		It("creates a file from the item", func() {
			client, err := NewStorage(dir, resync)
			Expect(err).NotTo(HaveOccurred())
			err = client.V1().Register()
			Expect(err).NotTo(HaveOccurred())
			schema := NewTestSchema1()
			createdSchema, err := client.V1().Schemas().Create(schema)
			Expect(err).NotTo(HaveOccurred())
			schema.Metadata = createdSchema.GetMetadata()
			Expect(schema).To(Equal(createdSchema))
		})
	})
	Describe("Create2Update", func() {
		It("creates and updates", func() {
			client, err := NewStorage(dir, resync)
			Expect(err).NotTo(HaveOccurred())
			err = client.V1().Register()
			Expect(err).NotTo(HaveOccurred())
			upstream := NewTestUpstream1()
			upstream, err = client.V1().Upstreams().Create(upstream)
			upstream2 := NewTestUpstream2()
			upstream2, err = client.V1().Upstreams().Create(upstream2)
			Expect(err).NotTo(HaveOccurred())

			_, err = client.V1().Upstreams().Update(upstream2)
			Expect(err).NotTo(HaveOccurred())

			created1, err := client.V1().Upstreams().Get(upstream.Name)
			Expect(err).NotTo(HaveOccurred())
			upstream.Metadata = created1.Metadata
			Expect(created1).To(Equal(upstream))

			created2, err := client.V1().Upstreams().Get(upstream2.Name)
			Expect(err).NotTo(HaveOccurred())
			upstream2.Metadata = created2.Metadata
			Expect(created2).To(Equal(upstream2))

		})
	})
	Describe("Create2Update vService", func() {
		It("creates and updates", func() {
			client, err := NewStorage(dir, resync)
			Expect(err).NotTo(HaveOccurred())
			err = client.V1().Register()
			Expect(err).NotTo(HaveOccurred())
			vService := NewTestVirtualService("v1")
			vService, err = client.V1().VirtualServices().Create(vService)
			vService2 := NewTestVirtualService("v2")
			vService2, err = client.V1().VirtualServices().Create(vService2)
			Expect(err).NotTo(HaveOccurred())

			_, err = client.V1().VirtualServices().Update(vService)
			Expect(err).NotTo(HaveOccurred())

			created1, err := client.V1().VirtualServices().Get(vService.Name)
			Expect(err).NotTo(HaveOccurred())
			vService.Metadata = created1.Metadata
			Expect(created1).To(Equal(vService))

			created2, err := client.V1().VirtualServices().Get(vService2.Name)
			Expect(err).NotTo(HaveOccurred())
			vService2.Metadata = created2.Metadata
			Expect(created2).To(Equal(vService2))
		})
	})
	Describe("Create2Update role", func() {
		It("creates and updates", func() {
			client, err := NewStorage(dir, resync)
			Expect(err).NotTo(HaveOccurred())
			err = client.V1().Register()
			Expect(err).NotTo(HaveOccurred())
			role := NewTestRole("v1")
			role, err = client.V1().Roles().Create(role)
			role2 := NewTestRole("v2")
			role2, err = client.V1().Roles().Create(role2)
			Expect(err).NotTo(HaveOccurred())

			_, err = client.V1().Roles().Update(role)
			Expect(err).NotTo(HaveOccurred())

			created1, err := client.V1().Roles().Get(role.Name)
			Expect(err).NotTo(HaveOccurred())
			role.Metadata = created1.Metadata
			Expect(created1).To(Equal(role))

			created2, err := client.V1().Roles().Get(role2.Name)
			Expect(err).NotTo(HaveOccurred())
			role2.Metadata = created2.Metadata
			Expect(created2).To(Equal(role2))
		})
	})
	Describe("Get", func() {
		It("gets a file from the name", func() {
			client, err := NewStorage(dir, resync)
			Expect(err).NotTo(HaveOccurred())
			err = client.V1().Register()
			Expect(err).NotTo(HaveOccurred())
			upstream := NewTestUpstream1()
			_, err = client.V1().Upstreams().Create(upstream)
			Expect(err).NotTo(HaveOccurred())
			created, err := client.V1().Upstreams().Get(upstream.Name)
			Expect(err).NotTo(HaveOccurred())
			upstream.Metadata = created.Metadata
			Expect(created).To(Equal(upstream))
		})
	})
	Describe("Update", func() {
		It("updates a file from the item", func() {
			client, err := NewStorage(dir, resync)
			Expect(err).NotTo(HaveOccurred())
			err = client.V1().Register()
			Expect(err).NotTo(HaveOccurred())
			upstream := NewTestUpstream1()
			created, err := client.V1().Upstreams().Create(upstream)
			Expect(err).NotTo(HaveOccurred())
			upstream.Type = "something-else"
			_, err = client.V1().Upstreams().Update(upstream)
			// need to set resource ver
			Expect(err).To(HaveOccurred())
			upstream.Metadata = created.GetMetadata()
			updated, err := client.V1().Upstreams().Update(upstream)
			Expect(err).NotTo(HaveOccurred())
			upstream.Metadata = updated.GetMetadata()
			Expect(updated).To(Equal(upstream))
		})
	})
	Describe("Delete", func() {
		It("deletes a file from the name", func() {
			client, err := NewStorage(dir, resync)
			Expect(err).NotTo(HaveOccurred())
			err = client.V1().Register()
			Expect(err).NotTo(HaveOccurred())
			upstream := NewTestUpstream1()
			_, err = client.V1().Upstreams().Create(upstream)
			Expect(err).NotTo(HaveOccurred())
			err = client.V1().Upstreams().Delete(upstream.Name)
			Expect(err).NotTo(HaveOccurred())
			_, err = client.V1().Upstreams().Get(upstream.Name)
			Expect(err).To(HaveOccurred())
		})
	})
})

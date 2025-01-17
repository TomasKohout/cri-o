package conmonmgr

import (
	runnerMock "github.com/cri-o/cri-o/test/mocks/cmdrunner"
	"github.com/cri-o/cri-o/utils/cmdrunner"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

const validPath = "/bin/ls"

// The actual test suite
var _ = t.Describe("ConmonManager", func() {
	var runner *runnerMock.MockCommandRunner
	t.Describe("New", func() {
		BeforeEach(func() {
			runner = runnerMock.NewMockCommandRunner(mockCtrl)
			cmdrunner.SetMocked(runner)
		})
		It("should fail when path not absolute", func() {
			// Given
			gomock.InOrder(
				runner.EXPECT().CombinedOutput(gomock.Any(), gomock.Any()).Return([]byte{}, nil),
			)

			// When
			mgr, err := New("")

			// Then
			Expect(err).NotTo(BeNil())
			Expect(mgr).To(BeNil())
		})
		It("should fail when command fails", func() {
			// Given
			gomock.InOrder(
				runner.EXPECT().CombinedOutput(gomock.Any(), gomock.Any()).Return([]byte{}, errors.New("cmd failed")),
			)

			// When
			mgr, err := New(validPath)

			// Then
			Expect(err).NotTo(BeNil())
			Expect(mgr).To(BeNil())
		})
		It("should fail when output unexpected", func() {
			// Given
			gomock.InOrder(
				runner.EXPECT().CombinedOutput(gomock.Any(), gomock.Any()).Return([]byte("unexpected"), nil),
			)

			// When
			mgr, err := New(validPath)

			// Then
			Expect(err).NotTo(BeNil())
			Expect(mgr).To(BeNil())
		})
		It("should succeed when output expected", func() {
			// Given
			gomock.InOrder(
				runner.EXPECT().CombinedOutput(gomock.Any(), gomock.Any()).Return([]byte("conmon version 2.0.0"), nil),
			)

			// When
			mgr, err := New(validPath)

			// Then
			Expect(err).To(BeNil())
			Expect(mgr).ToNot(BeNil())
		})
		It("should succeed when output expected", func() {
			// Given
			gomock.InOrder(
				runner.EXPECT().CombinedOutput(gomock.Any(), gomock.Any()).Return([]byte("conmon version 2.0.0"), nil),
			)

			// When
			mgr, err := New(validPath)

			// Then
			Expect(err).To(BeNil())
			Expect(mgr).ToNot(BeNil())
		})
	})
	t.Describe("parseConmonVersion", func() {
		var mgr *ConmonManager
		BeforeEach(func() {
			mgr = new(ConmonManager)
		})
		It("should fail when not number", func() {
			// When
			err := mgr.parseConmonVersion("invalid.0.0")
			// Then
			Expect(err).NotTo(BeNil())
		})
		It("should succeed when all are numbers", func() {
			// When
			err := mgr.parseConmonVersion("0.0.0")
			// Then
			Expect(err).To(BeNil())
		})
	})
	t.Describe("initializeSupportsSync", func() {
		var mgr *ConmonManager
		BeforeEach(func() {
			mgr = new(ConmonManager)
		})
		It("should be false when major version less", func() {
			// Given
			err := mgr.parseConmonVersion("1.0.19")
			Expect(err).To(BeNil())
			// When
			mgr.initializeSupportsSync()

			// Then
			Expect(mgr.SupportsSync()).To(Equal(false))
		})
		It("should be true when major version greater", func() {
			// Given
			err := mgr.parseConmonVersion("3.0.19")
			Expect(err).To(BeNil())

			// When
			mgr.initializeSupportsSync()

			// Then
			Expect(mgr.SupportsSync()).To(Equal(true))
		})
		It("should be true when minor version greater", func() {
			// Given
			err := mgr.parseConmonVersion("2.1.18")
			Expect(err).To(BeNil())

			// When
			mgr.initializeSupportsSync()

			// Then
			Expect(mgr.SupportsSync()).To(Equal(true))
		})
		It("should be false when patch version less", func() {
			// Given
			err := mgr.parseConmonVersion("2.0.18")
			Expect(err).To(BeNil())

			// When
			mgr.initializeSupportsSync()

			// Then
			Expect(mgr.SupportsSync()).To(Equal(false))
		})
		It("should be true when patch version greater", func() {
			// Given
			err := mgr.parseConmonVersion("2.0.20")
			Expect(err).To(BeNil())

			// When
			mgr.initializeSupportsSync()

			// Then
			Expect(mgr.SupportsSync()).To(Equal(true))
		})
		It("should be true when version equal", func() {
			// Given
			err := mgr.parseConmonVersion("2.0.19")
			Expect(err).To(BeNil())

			// When
			mgr.initializeSupportsSync()

			// Then
			Expect(mgr.SupportsSync()).To(Equal(true))
		})
	})
})

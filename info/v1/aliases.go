// Copyright 2026 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0.

// Package v1 is a thin compatibility alias façade over the canonical model
// package; the type DEFINITIONS now live in github.com/google/cadvisor/model.
// (collapse step 2b — design §6.2). Consumers may repoint to model directly;
// this façade lets them do so incrementally.
package v1

import model "github.com/google/cadvisor/model"

type AcceleratorStats = model.AcceleratorStats
type Cache = model.Cache
type CacheStats = model.CacheStats
type CloudProvider = model.CloudProvider
type ContainerDeletionEventData = model.ContainerDeletionEventData
type ContainerInfo = model.ContainerInfo
type ContainerInfoRequest = model.ContainerInfoRequest
type ContainerReference = model.ContainerReference
type ContainerReferenceSlice = model.ContainerReferenceSlice
type ContainerSpec = model.ContainerSpec
type ContainerStats = model.ContainerStats
type Core = model.Core
type CpuCFS = model.CpuCFS
type CpuSchedstat = model.CpuSchedstat
type CPUSetStats = model.CPUSetStats
type CpuSpec = model.CpuSpec
type CpuStats = model.CpuStats
type CpuUsage = model.CpuUsage
type DataType = model.DataType
type DiskInfo = model.DiskInfo
type DiskIoStats = model.DiskIoStats
type Event = model.Event
type EventData = model.EventData
type EventType = model.EventType
type FsInfo = model.FsInfo
type FsStats = model.FsStats
type Health = model.Health
type HugePagesInfo = model.HugePagesInfo
type HugetlbStats = model.HugetlbStats
type InstanceID = model.InstanceID
type InstanceType = model.InstanceType
type InterfaceStats = model.InterfaceStats
type LoadStats = model.LoadStats
type MachineInfo = model.MachineInfo
type MachineInfoFactory = model.MachineInfoFactory
type MemoryBandwidthStats = model.MemoryBandwidthStats
type MemoryEvents = model.MemoryEvents
type MemoryInfo = model.MemoryInfo
type MemoryNumaStats = model.MemoryNumaStats
type MemorySpec = model.MemorySpec
type MemoryStats = model.MemoryStats
type MemoryStatsMemoryData = model.MemoryStatsMemoryData
type MetricSpec = model.MetricSpec
type MetricType = model.MetricType
type MetricVal = model.MetricVal
type MetricValBasic = model.MetricValBasic
type NetInfo = model.NetInfo
type NetworkStats = model.NetworkStats
type Node = model.Node
type NVMInfo = model.NVMInfo
type OomKillEventData = model.OomKillEventData
type PerDiskStats = model.PerDiskStats
type PerfStat = model.PerfStat
type PerfUncoreStat = model.PerfUncoreStat
type PerfValue = model.PerfValue
type ProcessSpec = model.ProcessSpec
type ProcessStats = model.ProcessStats
type PSIData = model.PSIData
type PSIStats = model.PSIStats
type ResctrlStats = model.ResctrlStats
type TcpAdvancedStat = model.TcpAdvancedStat
type TcpStat = model.TcpStat
type UdpStat = model.UdpStat
type UlimitSpec = model.UlimitSpec
type VersionInfo = model.VersionInfo

const AWS = model.AWS
const Azure = model.Azure
const EventContainerCreation = model.EventContainerCreation
const EventContainerDeletion = model.EventContainerDeletion
const EventOom = model.EventOom
const EventOomKill = model.EventOomKill
const FloatType = model.FloatType
const GCE = model.GCE
const IntType = model.IntType
const MetricCumulative = model.MetricCumulative
const MetricGauge = model.MetricGauge
const UnknownInstance = model.UnknownInstance
const UnknownProvider = model.UnknownProvider
const UnNamedInstance = model.UnNamedInstance

var DefaultContainerInfoRequest = model.DefaultContainerInfoRequest

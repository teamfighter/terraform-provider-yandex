// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

func (m *PostgresqlConfig11_1C) SetMaxConnections(v *wrappers.Int64Value) {
	m.MaxConnections = v
}

func (m *PostgresqlConfig11_1C) SetSharedBuffers(v *wrappers.Int64Value) {
	m.SharedBuffers = v
}

func (m *PostgresqlConfig11_1C) SetTempBuffers(v *wrappers.Int64Value) {
	m.TempBuffers = v
}

func (m *PostgresqlConfig11_1C) SetMaxPreparedTransactions(v *wrappers.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *PostgresqlConfig11_1C) SetWorkMem(v *wrappers.Int64Value) {
	m.WorkMem = v
}

func (m *PostgresqlConfig11_1C) SetMaintenanceWorkMem(v *wrappers.Int64Value) {
	m.MaintenanceWorkMem = v
}

func (m *PostgresqlConfig11_1C) SetAutovacuumWorkMem(v *wrappers.Int64Value) {
	m.AutovacuumWorkMem = v
}

func (m *PostgresqlConfig11_1C) SetTempFileLimit(v *wrappers.Int64Value) {
	m.TempFileLimit = v
}

func (m *PostgresqlConfig11_1C) SetVacuumCostDelay(v *wrappers.Int64Value) {
	m.VacuumCostDelay = v
}

func (m *PostgresqlConfig11_1C) SetVacuumCostPageHit(v *wrappers.Int64Value) {
	m.VacuumCostPageHit = v
}

func (m *PostgresqlConfig11_1C) SetVacuumCostPageMiss(v *wrappers.Int64Value) {
	m.VacuumCostPageMiss = v
}

func (m *PostgresqlConfig11_1C) SetVacuumCostPageDirty(v *wrappers.Int64Value) {
	m.VacuumCostPageDirty = v
}

func (m *PostgresqlConfig11_1C) SetVacuumCostLimit(v *wrappers.Int64Value) {
	m.VacuumCostLimit = v
}

func (m *PostgresqlConfig11_1C) SetBgwriterDelay(v *wrappers.Int64Value) {
	m.BgwriterDelay = v
}

func (m *PostgresqlConfig11_1C) SetBgwriterLruMaxpages(v *wrappers.Int64Value) {
	m.BgwriterLruMaxpages = v
}

func (m *PostgresqlConfig11_1C) SetBgwriterLruMultiplier(v *wrappers.DoubleValue) {
	m.BgwriterLruMultiplier = v
}

func (m *PostgresqlConfig11_1C) SetBgwriterFlushAfter(v *wrappers.Int64Value) {
	m.BgwriterFlushAfter = v
}

func (m *PostgresqlConfig11_1C) SetBackendFlushAfter(v *wrappers.Int64Value) {
	m.BackendFlushAfter = v
}

func (m *PostgresqlConfig11_1C) SetOldSnapshotThreshold(v *wrappers.Int64Value) {
	m.OldSnapshotThreshold = v
}

func (m *PostgresqlConfig11_1C) SetWalLevel(v PostgresqlConfig11_1C_WalLevel) {
	m.WalLevel = v
}

func (m *PostgresqlConfig11_1C) SetSynchronousCommit(v PostgresqlConfig11_1C_SynchronousCommit) {
	m.SynchronousCommit = v
}

func (m *PostgresqlConfig11_1C) SetCheckpointTimeout(v *wrappers.Int64Value) {
	m.CheckpointTimeout = v
}

func (m *PostgresqlConfig11_1C) SetCheckpointCompletionTarget(v *wrappers.DoubleValue) {
	m.CheckpointCompletionTarget = v
}

func (m *PostgresqlConfig11_1C) SetCheckpointFlushAfter(v *wrappers.Int64Value) {
	m.CheckpointFlushAfter = v
}

func (m *PostgresqlConfig11_1C) SetMaxWalSize(v *wrappers.Int64Value) {
	m.MaxWalSize = v
}

func (m *PostgresqlConfig11_1C) SetMinWalSize(v *wrappers.Int64Value) {
	m.MinWalSize = v
}

func (m *PostgresqlConfig11_1C) SetMaxStandbyStreamingDelay(v *wrappers.Int64Value) {
	m.MaxStandbyStreamingDelay = v
}

func (m *PostgresqlConfig11_1C) SetDefaultStatisticsTarget(v *wrappers.Int64Value) {
	m.DefaultStatisticsTarget = v
}

func (m *PostgresqlConfig11_1C) SetConstraintExclusion(v PostgresqlConfig11_1C_ConstraintExclusion) {
	m.ConstraintExclusion = v
}

func (m *PostgresqlConfig11_1C) SetCursorTupleFraction(v *wrappers.DoubleValue) {
	m.CursorTupleFraction = v
}

func (m *PostgresqlConfig11_1C) SetFromCollapseLimit(v *wrappers.Int64Value) {
	m.FromCollapseLimit = v
}

func (m *PostgresqlConfig11_1C) SetJoinCollapseLimit(v *wrappers.Int64Value) {
	m.JoinCollapseLimit = v
}

func (m *PostgresqlConfig11_1C) SetForceParallelMode(v PostgresqlConfig11_1C_ForceParallelMode) {
	m.ForceParallelMode = v
}

func (m *PostgresqlConfig11_1C) SetClientMinMessages(v PostgresqlConfig11_1C_LogLevel) {
	m.ClientMinMessages = v
}

func (m *PostgresqlConfig11_1C) SetLogMinMessages(v PostgresqlConfig11_1C_LogLevel) {
	m.LogMinMessages = v
}

func (m *PostgresqlConfig11_1C) SetLogMinErrorStatement(v PostgresqlConfig11_1C_LogLevel) {
	m.LogMinErrorStatement = v
}

func (m *PostgresqlConfig11_1C) SetLogMinDurationStatement(v *wrappers.Int64Value) {
	m.LogMinDurationStatement = v
}

func (m *PostgresqlConfig11_1C) SetLogCheckpoints(v *wrappers.BoolValue) {
	m.LogCheckpoints = v
}

func (m *PostgresqlConfig11_1C) SetLogConnections(v *wrappers.BoolValue) {
	m.LogConnections = v
}

func (m *PostgresqlConfig11_1C) SetLogDisconnections(v *wrappers.BoolValue) {
	m.LogDisconnections = v
}

func (m *PostgresqlConfig11_1C) SetLogDuration(v *wrappers.BoolValue) {
	m.LogDuration = v
}

func (m *PostgresqlConfig11_1C) SetLogErrorVerbosity(v PostgresqlConfig11_1C_LogErrorVerbosity) {
	m.LogErrorVerbosity = v
}

func (m *PostgresqlConfig11_1C) SetLogLockWaits(v *wrappers.BoolValue) {
	m.LogLockWaits = v
}

func (m *PostgresqlConfig11_1C) SetLogStatement(v PostgresqlConfig11_1C_LogStatement) {
	m.LogStatement = v
}

func (m *PostgresqlConfig11_1C) SetLogTempFiles(v *wrappers.Int64Value) {
	m.LogTempFiles = v
}

func (m *PostgresqlConfig11_1C) SetSearchPath(v string) {
	m.SearchPath = v
}

func (m *PostgresqlConfig11_1C) SetRowSecurity(v *wrappers.BoolValue) {
	m.RowSecurity = v
}

func (m *PostgresqlConfig11_1C) SetDefaultTransactionIsolation(v PostgresqlConfig11_1C_TransactionIsolation) {
	m.DefaultTransactionIsolation = v
}

func (m *PostgresqlConfig11_1C) SetStatementTimeout(v *wrappers.Int64Value) {
	m.StatementTimeout = v
}

func (m *PostgresqlConfig11_1C) SetLockTimeout(v *wrappers.Int64Value) {
	m.LockTimeout = v
}

func (m *PostgresqlConfig11_1C) SetIdleInTransactionSessionTimeout(v *wrappers.Int64Value) {
	m.IdleInTransactionSessionTimeout = v
}

func (m *PostgresqlConfig11_1C) SetByteaOutput(v PostgresqlConfig11_1C_ByteaOutput) {
	m.ByteaOutput = v
}

func (m *PostgresqlConfig11_1C) SetXmlbinary(v PostgresqlConfig11_1C_XmlBinary) {
	m.Xmlbinary = v
}

func (m *PostgresqlConfig11_1C) SetXmloption(v PostgresqlConfig11_1C_XmlOption) {
	m.Xmloption = v
}

func (m *PostgresqlConfig11_1C) SetGinPendingListLimit(v *wrappers.Int64Value) {
	m.GinPendingListLimit = v
}

func (m *PostgresqlConfig11_1C) SetDeadlockTimeout(v *wrappers.Int64Value) {
	m.DeadlockTimeout = v
}

func (m *PostgresqlConfig11_1C) SetMaxLocksPerTransaction(v *wrappers.Int64Value) {
	m.MaxLocksPerTransaction = v
}

func (m *PostgresqlConfig11_1C) SetMaxPredLocksPerTransaction(v *wrappers.Int64Value) {
	m.MaxPredLocksPerTransaction = v
}

func (m *PostgresqlConfig11_1C) SetArrayNulls(v *wrappers.BoolValue) {
	m.ArrayNulls = v
}

func (m *PostgresqlConfig11_1C) SetBackslashQuote(v PostgresqlConfig11_1C_BackslashQuote) {
	m.BackslashQuote = v
}

func (m *PostgresqlConfig11_1C) SetDefaultWithOids(v *wrappers.BoolValue) {
	m.DefaultWithOids = v
}

func (m *PostgresqlConfig11_1C) SetEscapeStringWarning(v *wrappers.BoolValue) {
	m.EscapeStringWarning = v
}

func (m *PostgresqlConfig11_1C) SetLoCompatPrivileges(v *wrappers.BoolValue) {
	m.LoCompatPrivileges = v
}

func (m *PostgresqlConfig11_1C) SetOperatorPrecedenceWarning(v *wrappers.BoolValue) {
	m.OperatorPrecedenceWarning = v
}

func (m *PostgresqlConfig11_1C) SetQuoteAllIdentifiers(v *wrappers.BoolValue) {
	m.QuoteAllIdentifiers = v
}

func (m *PostgresqlConfig11_1C) SetStandardConformingStrings(v *wrappers.BoolValue) {
	m.StandardConformingStrings = v
}

func (m *PostgresqlConfig11_1C) SetSynchronizeSeqscans(v *wrappers.BoolValue) {
	m.SynchronizeSeqscans = v
}

func (m *PostgresqlConfig11_1C) SetTransformNullEquals(v *wrappers.BoolValue) {
	m.TransformNullEquals = v
}

func (m *PostgresqlConfig11_1C) SetExitOnError(v *wrappers.BoolValue) {
	m.ExitOnError = v
}

func (m *PostgresqlConfig11_1C) SetSeqPageCost(v *wrappers.DoubleValue) {
	m.SeqPageCost = v
}

func (m *PostgresqlConfig11_1C) SetRandomPageCost(v *wrappers.DoubleValue) {
	m.RandomPageCost = v
}

func (m *PostgresqlConfig11_1C) SetAutovacuumMaxWorkers(v *wrappers.Int64Value) {
	m.AutovacuumMaxWorkers = v
}

func (m *PostgresqlConfig11_1C) SetAutovacuumVacuumCostDelay(v *wrappers.Int64Value) {
	m.AutovacuumVacuumCostDelay = v
}

func (m *PostgresqlConfig11_1C) SetAutovacuumVacuumCostLimit(v *wrappers.Int64Value) {
	m.AutovacuumVacuumCostLimit = v
}

func (m *PostgresqlConfig11_1C) SetAutovacuumNaptime(v *wrappers.Int64Value) {
	m.AutovacuumNaptime = v
}

func (m *PostgresqlConfig11_1C) SetArchiveTimeout(v *wrappers.Int64Value) {
	m.ArchiveTimeout = v
}

func (m *PostgresqlConfig11_1C) SetTrackActivityQuerySize(v *wrappers.Int64Value) {
	m.TrackActivityQuerySize = v
}

func (m *PostgresqlConfig11_1C) SetEnableBitmapscan(v *wrappers.BoolValue) {
	m.EnableBitmapscan = v
}

func (m *PostgresqlConfig11_1C) SetEnableHashagg(v *wrappers.BoolValue) {
	m.EnableHashagg = v
}

func (m *PostgresqlConfig11_1C) SetEnableHashjoin(v *wrappers.BoolValue) {
	m.EnableHashjoin = v
}

func (m *PostgresqlConfig11_1C) SetEnableIndexscan(v *wrappers.BoolValue) {
	m.EnableIndexscan = v
}

func (m *PostgresqlConfig11_1C) SetEnableIndexonlyscan(v *wrappers.BoolValue) {
	m.EnableIndexonlyscan = v
}

func (m *PostgresqlConfig11_1C) SetEnableMaterial(v *wrappers.BoolValue) {
	m.EnableMaterial = v
}

func (m *PostgresqlConfig11_1C) SetEnableMergejoin(v *wrappers.BoolValue) {
	m.EnableMergejoin = v
}

func (m *PostgresqlConfig11_1C) SetEnableNestloop(v *wrappers.BoolValue) {
	m.EnableNestloop = v
}

func (m *PostgresqlConfig11_1C) SetEnableSeqscan(v *wrappers.BoolValue) {
	m.EnableSeqscan = v
}

func (m *PostgresqlConfig11_1C) SetEnableSort(v *wrappers.BoolValue) {
	m.EnableSort = v
}

func (m *PostgresqlConfig11_1C) SetEnableTidscan(v *wrappers.BoolValue) {
	m.EnableTidscan = v
}

func (m *PostgresqlConfig11_1C) SetMaxWorkerProcesses(v *wrappers.Int64Value) {
	m.MaxWorkerProcesses = v
}

func (m *PostgresqlConfig11_1C) SetMaxParallelWorkers(v *wrappers.Int64Value) {
	m.MaxParallelWorkers = v
}

func (m *PostgresqlConfig11_1C) SetMaxParallelWorkersPerGather(v *wrappers.Int64Value) {
	m.MaxParallelWorkersPerGather = v
}

func (m *PostgresqlConfig11_1C) SetAutovacuumVacuumScaleFactor(v *wrappers.DoubleValue) {
	m.AutovacuumVacuumScaleFactor = v
}

func (m *PostgresqlConfig11_1C) SetAutovacuumAnalyzeScaleFactor(v *wrappers.DoubleValue) {
	m.AutovacuumAnalyzeScaleFactor = v
}

func (m *PostgresqlConfig11_1C) SetDefaultTransactionReadOnly(v *wrappers.BoolValue) {
	m.DefaultTransactionReadOnly = v
}

func (m *PostgresqlConfig11_1C) SetTimezone(v string) {
	m.Timezone = v
}

func (m *PostgresqlConfig11_1C) SetEnableParallelAppend(v *wrappers.BoolValue) {
	m.EnableParallelAppend = v
}

func (m *PostgresqlConfig11_1C) SetEnableParallelHash(v *wrappers.BoolValue) {
	m.EnableParallelHash = v
}

func (m *PostgresqlConfig11_1C) SetEnablePartitionPruning(v *wrappers.BoolValue) {
	m.EnablePartitionPruning = v
}

func (m *PostgresqlConfig11_1C) SetEnablePartitionwiseAggregate(v *wrappers.BoolValue) {
	m.EnablePartitionwiseAggregate = v
}

func (m *PostgresqlConfig11_1C) SetEnablePartitionwiseJoin(v *wrappers.BoolValue) {
	m.EnablePartitionwiseJoin = v
}

func (m *PostgresqlConfig11_1C) SetJit(v *wrappers.BoolValue) {
	m.Jit = v
}

func (m *PostgresqlConfig11_1C) SetMaxParallelMaintenanceWorkers(v *wrappers.Int64Value) {
	m.MaxParallelMaintenanceWorkers = v
}

func (m *PostgresqlConfig11_1C) SetParallelLeaderParticipation(v *wrappers.BoolValue) {
	m.ParallelLeaderParticipation = v
}

func (m *PostgresqlConfig11_1C) SetVacuumCleanupIndexScaleFactor(v *wrappers.DoubleValue) {
	m.VacuumCleanupIndexScaleFactor = v
}

func (m *PostgresqlConfig11_1C) SetEffectiveIoConcurrency(v *wrappers.Int64Value) {
	m.EffectiveIoConcurrency = v
}

func (m *PostgresqlConfig11_1C) SetEffectiveCacheSize(v *wrappers.Int64Value) {
	m.EffectiveCacheSize = v
}

func (m *PostgresqlConfig11_1C) SetSharedPreloadLibraries(v []PostgresqlConfig11_1C_SharedPreloadLibraries) {
	m.SharedPreloadLibraries = v
}

func (m *PostgresqlConfig11_1C) SetAutoExplainLogMinDuration(v *wrappers.Int64Value) {
	m.AutoExplainLogMinDuration = v
}

func (m *PostgresqlConfig11_1C) SetAutoExplainLogAnalyze(v *wrappers.BoolValue) {
	m.AutoExplainLogAnalyze = v
}

func (m *PostgresqlConfig11_1C) SetAutoExplainLogBuffers(v *wrappers.BoolValue) {
	m.AutoExplainLogBuffers = v
}

func (m *PostgresqlConfig11_1C) SetAutoExplainLogTiming(v *wrappers.BoolValue) {
	m.AutoExplainLogTiming = v
}

func (m *PostgresqlConfig11_1C) SetAutoExplainLogTriggers(v *wrappers.BoolValue) {
	m.AutoExplainLogTriggers = v
}

func (m *PostgresqlConfig11_1C) SetAutoExplainLogVerbose(v *wrappers.BoolValue) {
	m.AutoExplainLogVerbose = v
}

func (m *PostgresqlConfig11_1C) SetAutoExplainLogNestedStatements(v *wrappers.BoolValue) {
	m.AutoExplainLogNestedStatements = v
}

func (m *PostgresqlConfig11_1C) SetAutoExplainSampleRate(v *wrappers.DoubleValue) {
	m.AutoExplainSampleRate = v
}

func (m *PostgresqlConfig11_1C) SetPgHintPlanEnableHint(v *wrappers.BoolValue) {
	m.PgHintPlanEnableHint = v
}

func (m *PostgresqlConfig11_1C) SetPgHintPlanEnableHintTable(v *wrappers.BoolValue) {
	m.PgHintPlanEnableHintTable = v
}

func (m *PostgresqlConfig11_1C) SetPgHintPlanDebugPrint(v PostgresqlConfig11_1C_PgHintPlanDebugPrint) {
	m.PgHintPlanDebugPrint = v
}

func (m *PostgresqlConfig11_1C) SetPgHintPlanMessageLevel(v PostgresqlConfig11_1C_LogLevel) {
	m.PgHintPlanMessageLevel = v
}

func (m *PostgresqlConfigSet11_1C) SetEffectiveConfig(v *PostgresqlConfig11_1C) {
	m.EffectiveConfig = v
}

func (m *PostgresqlConfigSet11_1C) SetUserConfig(v *PostgresqlConfig11_1C) {
	m.UserConfig = v
}

func (m *PostgresqlConfigSet11_1C) SetDefaultConfig(v *PostgresqlConfig11_1C) {
	m.DefaultConfig = v
}
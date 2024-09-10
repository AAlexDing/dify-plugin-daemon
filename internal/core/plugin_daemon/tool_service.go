package plugin_daemon

import (
	"github.com/langgenius/dify-plugin-daemon/internal/core/session_manager"
	"github.com/langgenius/dify-plugin-daemon/internal/types/entities/requests"
	"github.com/langgenius/dify-plugin-daemon/internal/types/entities/tool_entities"
	"github.com/langgenius/dify-plugin-daemon/internal/utils/stream"
)

func InvokeTool(
	session *session_manager.Session,
	request *requests.RequestInvokeTool,
) (
	*stream.Stream[tool_entities.ToolResponseChunk], error,
) {
	return genericInvokePlugin[requests.RequestInvokeTool, tool_entities.ToolResponseChunk](
		session,
		request,
		128,
	)
}

func ValidateToolCredentials(
	session *session_manager.Session,
	request *requests.RequestValidateToolCredentials,
) (
	*stream.Stream[tool_entities.ValidateCredentialsResult], error,
) {
	return genericInvokePlugin[requests.RequestValidateToolCredentials, tool_entities.ValidateCredentialsResult](
		session,
		request,
		1,
	)
}

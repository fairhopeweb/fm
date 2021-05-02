package app

import (
	"fmt"

	"github.com/knipferrc/fm/components"
	"github.com/knipferrc/fm/constants"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) getLeftPane() string {
	border := lipgloss.NormalBorder()
	borderRightWidth := lipgloss.Width(border.Right + border.Top)
	borderLeftWidth := lipgloss.Width(border.Left + border.Top)
	halfScreenWidth := m.ScreenWidth / 2
	leftPaneActive := m.ActivePane == constants.PrimaryPane
	paneHeight := m.ScreenHeight - constants.StatusBarHeight - borderRightWidth

	return components.Pane(
		halfScreenWidth-borderLeftWidth,
		paneHeight,
		leftPaneActive,
		m.PrimaryViewport.View(),
	)
}

func (m Model) getRightPane() string {
	border := lipgloss.NormalBorder()
	borderRightWidth := lipgloss.Width(border.Right + border.Top)
	halfScreenWidth := m.ScreenWidth / 2
	rightPaneActive := m.ActivePane == constants.SecondaryPane
	paneHeight := m.ScreenHeight - constants.StatusBarHeight - borderRightWidth
	rightPane := ""

	if m.ShowMovePrompt {
		rightPane = components.Pane(
			halfScreenWidth-borderRightWidth,
			paneHeight,
			rightPaneActive,
			components.MovePrompt(m.Textinput.View()),
		)
	} else if m.ShowRenamePrompt {
		rightPane = components.Pane(
			halfScreenWidth-borderRightWidth,
			paneHeight,
			rightPaneActive,
			components.RenamePrompt(m.Textinput.View()),
		)
	} else if m.ShowDeletePrompt {
		rightPane = components.Pane(
			halfScreenWidth-borderRightWidth,
			paneHeight,
			rightPaneActive,
			components.DeletePrompt(m.Textinput.View(), m.Files[m.Cursor].Name()),
		)
	} else if m.ShowCreatePrompt {
		rightPane = components.Pane(
			halfScreenWidth-borderRightWidth,
			paneHeight,
			rightPaneActive,
			components.CreatePrompt(m.Textinput.View()),
		)
	} else {
		rightPane = components.Pane(
			halfScreenWidth-borderRightWidth,
			paneHeight,
			rightPaneActive,
			m.SecondaryViewport.View(),
		)
	}

	return rightPane
}

func (m Model) View() string {
	if !m.Ready || len(m.Files) <= 0 {
		return fmt.Sprintf("%s%s", m.Spinner.View(), "loading...")
	}

	leftPane := m.getLeftPane()
	rightPane := m.getRightPane()

	panes := lipgloss.JoinHorizontal(lipgloss.Top, leftPane, rightPane)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		panes,
		components.StatusBar(m.ScreenWidth, m.Cursor, len(m.Files), m.Files[m.Cursor]),
	)
}
